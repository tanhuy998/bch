package refreshTokenService

import (
	"app/internal/bootstrap"
	"app/internal/common"
	"app/internal/generalToken"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	jwtTokenServicePort "app/port/jwtTokenService"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenBlackListServicePort "app/port/refreshTokenBlackList"
	refreshTokenIDService "app/port/refreshTokenID"
	"app/valueObject"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	exp_duration = time.Hour * 2
)

var (
	ERR_CANNOT_REVOKE_REFRESH_TOKEN = errors.New("RefreshTokenManipulatorService error: cannot revoke refresh token")
)

type (
	GeneralTokenID           = generalToken.GeneralTokenID
	IRefreshTokenManipulator = refreshTokenServicePort.IRefreshTokenManipulator
	IRefreshToken            = refreshTokenServicePort.IRefreshToken
	GeneralToken             = generalToken.GeneralTokenID

	ClaimsOption = func(claims *jwt.RegisteredClaims)

	RefreshTokenManipulatorService struct {
		AudienceList           accessTokenServicePort.AccessTokenAudienceList
		RefreshTokenIDProvider refreshTokenIDService.IRefreshTokenIDProvider
		JWTTokenService        jwtTokenServicePort.ISymmetricJWTTokenManipulator
		RefreshTokenBlackList  refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator
	}
)

func (this *RefreshTokenManipulatorService) Generate(
	tenantUUID uuid.UUID, generalToken GeneralToken, ctx context.Context,
) (refreshTokenServicePort.IRefreshToken, error) {

	return this.makeFor(tenantUUID, generalToken)
}

func (this *RefreshTokenManipulatorService) makeFor(
	tenantUUID uuid.UUID, generalToken GeneralToken, claimOption ...ClaimsOption,
) (refreshTokenServicePort.IRefreshToken, error) {

	token := this.JWTTokenService.GenerateToken()

	refreshTokenID, err := this.RefreshTokenIDProvider.Generate(generalToken)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	customClaims := &jwt_refresh_token_custom_claims{
		jwt.RegisteredClaims{
			Issuer: bootstrap.GetAppName(),
			//Audience:  jwt.ClaimStrings(bootstrap.GetHostNames()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp_duration)),
		},
		refreshTokenID,
		libCommon.PointerPrimitive(tenantUUID),
	}

	token.Claims = customClaims

	for _, optionFunc := range claimOption {

		optionFunc(&customClaims.RegisteredClaims)
	}

	return newFromToken(token)
}

func (this *RefreshTokenManipulatorService) Revoke(refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context) error {

	refreshTokenID := refreshToken.GetTokenID()

	payload := &valueObject.RefreshTokenBlackListPayload{
		UserUUID: libCommon.PointerPrimitive(refreshToken.GetUserUUID()),
	}

	exp, err := refreshToken.GetExpireTime()

	if errors.Is(err, common.ERR_INTERNAL) {

		return err
	}

	if err != nil {

		return libError.NewInternal(err)
	}

	err = this.RefreshTokenBlackList.SetWithExpire(refreshTokenID, payload, *exp, ctx)

	if err != nil {

		return err
	}

	return nil
}

func (this *RefreshTokenManipulatorService) SignString(refreshToken IRefreshToken) (string, error) {

	if v, ok := any(refreshToken).(*jwt_refresh_token); ok {

		return this.JWTTokenService.SignString(v.jwt_token)
	}

	return "", ERR_INVALID_TOKEN
}

func (this *RefreshTokenManipulatorService) Read(str string) (IRefreshToken, error) {

	token, err := this.JWTTokenService.VerifyTokenStringCustomClaim(str, &jwt_refresh_token_custom_claims{})

	if err != nil {

		return nil, err
	}

	return newFromToken(token)
}

func (this *RefreshTokenManipulatorService) DefaultExpireDuration() time.Duration {

	return exp_duration
}

func (this *RefreshTokenManipulatorService) Rotate(refreshToken IRefreshToken, ctx context.Context) (IRefreshToken, error) {

	switch {
	case refreshToken == nil:
		return nil, errors.New("nil passed to RefreshTokenManipulatorService.Rotate")
	case refreshToken.Expired():
		return nil, refreshTokenServicePort.ERR_TOKEN_EXPIRE
	}

	refreshTokenInBlackList, err := this.RefreshTokenBlackList.Has(refreshToken.GetTokenID(), ctx)

	switch {
	case err != nil:
		return nil, err
	case refreshTokenInBlackList:
		return nil, errors.New("refreshTokenManipulatorService error: refresh token in blacklist")
	}

	exp, err := refreshToken.GetExpireTime()

	if errors.Is(err, common.ERR_INTERNAL) {

		return nil, err
	}

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	generalID, _, err := this.RefreshTokenIDProvider.Extract(refreshToken.GetTokenID())

	if err != nil {

		return nil, err
	}

	if exp == nil {

		return this.makeFor(refreshToken.GetTenantUUID(), generalID)
	}

	err = this.Revoke(refreshToken, ctx)

	if err != nil {

		return nil, err
	}

	return this.makeFor(refreshToken.GetUserUUID(), generalID, SetExpire(*exp))
}
