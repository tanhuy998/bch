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
	"app/unitOfWork"
	jwtClaim "app/valueObject/jwt"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	default_exp_duration = time.Hour * 2
	test_exp_duration    = time.Minute * 5
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
		AudienceList accessTokenServicePort.AccessTokenAudienceList
		unitOfWork.OperationLogger
		RefreshTokenIDProvider refreshTokenIDService.IRefreshTokenIDProvider
		JWTTokenService        jwtTokenServicePort.ISymmetricJWTTokenManipulator
	}
)

func (this *RefreshTokenManipulatorService) Generate(
	tenantUUID uuid.UUID, generalToken GeneralToken, ctx context.Context,
) (refreshTokenServicePort.IRefreshToken, error) {

	return this.makeFor(tenantUUID, generalToken.GetTokenID(), ctx)
}

func (this *RefreshTokenManipulatorService) makeFor(
	tenantUUID uuid.UUID, generalTokenID generalToken.GeneralTokenID, ctx context.Context, claimOption ...ClaimsOption,
) (ret *jwt_refresh_token, err error) {

	defer this.OperationLogger.PushTraceCondWithMessurement("generate_refresh_token", ctx)("success", err, "")

	token := this.JWTTokenService.GenerateToken()

	refreshTokenID, err := this.RefreshTokenIDProvider.Generate(generalTokenID)

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	customClaims := &jwt_refresh_token_custom_claims{
		Issuer:         bootstrap.GetAppName(),
		IssuedAt:       jwt.NewNumericDate(time.Now()),
		RefreshTokenID: refreshTokenID,
		TenantUUID:     libCommon.PointerPrimitive(tenantUUID),
		ExpiresAt:      jwt.NewNumericDate(time.Now().Add(default_exp_duration)),
	}

	if bootstrap.IsTestingRefreshToken() {

		customClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(test_exp_duration))
	}

	jwtClaim.SetupRefreshToken(&customClaims.PrivateClaims)

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

	return default_exp_duration
}

func (this *RefreshTokenManipulatorService) makeBase(refreshToken IRefreshToken, ctx context.Context) (*jwt_refresh_token, error) {

	generalTokenID, _, err := this.RefreshTokenIDProvider.Extract(refreshToken.GetTokenID())

	if err != nil {

		return nil, libError.NewInternal(err)
	}

	return this.makeFor(refreshToken.GetTenantUUID(), generalTokenID, ctx)
}

func (this *RefreshTokenManipulatorService) Rotate(refreshToken IRefreshToken, ctx context.Context) (IRefreshToken, error) {

	switch {
	case refreshToken == nil:
		return nil, errors.New("nil passed to RefreshTokenManipulatorService.Rotate")
	case refreshToken.Expired():
		return nil, refreshTokenServicePort.ERR_TOKEN_EXPIRE
	}

	return this.makeBase(refreshToken, ctx)
}
