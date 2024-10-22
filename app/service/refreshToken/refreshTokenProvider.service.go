package refreshTokenService

import (
	"app/internal/bootstrap"
	"app/internal/generalToken"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"
	jwtTokenServicePort "app/port/jwtTokenService"
	refreshTokenServicePort "app/port/refreshToken"
	refreshTokenIDService "app/port/refreshTokenID"
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
	GeneralToken             = generalTokenServicePort.IGeneralToken

	ClaimsOption = func(claims *jwt.RegisteredClaims)

	RefreshTokenManipulatorService struct {
		AudienceList           accessTokenServicePort.AccessTokenAudienceList
		RefreshTokenIDProvider refreshTokenIDService.IRefreshTokenIDProvider
		JWTTokenService        jwtTokenServicePort.ISymmetricJWTTokenManipulator
		//RefreshTokenBlackList  refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator
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

	refreshTokenID, err := this.RefreshTokenIDProvider.Generate(generalToken.GetTokenID())

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	customClaims := &jwt_refresh_token_custom_claims{
		Issuer:         bootstrap.GetAppName(),
		IssuedAt:       jwt.NewNumericDate(time.Now()),
		RefreshTokenID: refreshTokenID,
		TenantUUID:     libCommon.PointerPrimitive(tenantUUID),
	}

	if generalToken.GetExpiretime() != nil {

		customClaims.ExpiresAt = jwt.NewNumericDate(*generalToken.GetExpiretime())
	}

	token.Claims = customClaims

	for _, optionFunc := range claimOption {

		optionFunc(&customClaims.RegisteredClaims)
	}

	rt, err := newFromToken(token, this.RefreshTokenIDProvider)

	if err != nil {

		return nil, err
	}

	// rt.userUUID = generalToken.GetUserUUID() // libCommon.PointerPrimitive(generalToken.GetUserUUID())

	return rt, nil
}

func (this *RefreshTokenManipulatorService) Revoke(refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context) error {

	// refreshTokenID := refreshToken.GetTokenID()

	// payload := &valueObject.RefreshTokenBlackListPayload{
	// 	UserUUID: libCommon.PointerPrimitive(refreshToken.GetUserUUID()),
	// }

	// exp, err := refreshToken.GetExpireTime()

	// if errors.Is(err, common.ERR_INTERNAL) {

	// 	return err
	// }

	// if err != nil {

	// 	return libError.NewInternal(err)
	// }

	// exp := refreshToken.GetExpireTime()

	// err := this.RefreshTokenBlackList.SetWithExpire(refreshTokenID, payload, *exp, ctx)

	// if err != nil {

	// 	return err
	// }

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

	return newFromToken(token, this.RefreshTokenIDProvider)
}

func (this *RefreshTokenManipulatorService) DefaultExpireDuration() time.Duration {

	return exp_duration
}

func (this *RefreshTokenManipulatorService) makeBase(refreshToken IRefreshToken) (IRefreshToken, error) {

	token := this.JWTTokenService.GenerateToken()

	generalTokenID, _, err := this.RefreshTokenIDProvider.Extract(refreshToken.GetTokenID())

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	newRefreshTokenID, err := this.RefreshTokenIDProvider.Generate(generalTokenID)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	customClaims := &jwt_refresh_token_custom_claims{
		Issuer:   bootstrap.GetAppName(),
		IssuedAt: jwt.NewNumericDate(time.Now()),

		RefreshTokenID: newRefreshTokenID,
		TenantUUID:     libCommon.PointerPrimitive(refreshToken.GetTenantUUID()),
	}

	if refreshToken.GetExpireTime() != nil {
		customClaims.ExpiresAt = jwt.NewNumericDate(*refreshToken.GetExpireTime())
	}

	token.Claims = customClaims

	return newFromToken(token, this.RefreshTokenIDProvider)
}

func (this *RefreshTokenManipulatorService) Rotate(refreshToken IRefreshToken, ctx context.Context) (IRefreshToken, error) {

	switch {
	case refreshToken == nil:
		return nil, errors.New("nil passed to RefreshTokenManipulatorService.Rotate")
	case refreshToken.Expired():
		return nil, refreshTokenServicePort.ERR_TOKEN_EXPIRE
	}

	// refreshTokenInBlackList, err := this.RefreshTokenBlackList.Has(refreshToken.GetTokenID(), ctx)

	// switch {
	// case err != nil:
	// 	return nil, err
	// case refreshTokenInBlackList:
	// 	return nil, libError.NewInternal(fmt.Errorf("refresh token in black list"))
	// }

	// exp, err := refreshToken.GetExpireTime()

	// if errors.Is(err, common.ERR_INTERNAL) {

	// 	return nil, err
	// }

	// if err != nil {

	// 	return nil, libError.NewInternal(err)
	// }

	// exp := refreshToken.GetExpireTime()

	// generalID, _, err := this.RefreshTokenIDProvider.Extract(refreshToken.GetTokenID())

	// if err != nil {

	// 	return nil, err
	// }

	// if exp == nil {

	// 	// return this.makeFor(refreshToken.GetTenantUUID(), generalID)
	// 	return this.makeBase(refreshToken)
	// }

	// err = this.Revoke(refreshToken, ctx)

	// if err != nil {

	// 	return nil, err
	// }

	// return this.makeFor(refreshToken.GetUserUUID(), generalID, SetExpire(*exp))

	return this.makeBase(refreshToken)
}
