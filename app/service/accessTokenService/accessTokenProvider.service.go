package accessTokenService

import (
	"app/internal/bootstrap"
	"app/internal/common"
	libCommon "app/internal/lib/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	authServicePort "app/port/auth"
	generalTokenServicePort "app/port/generalToken"
	jwtTokenServicePort "app/port/jwtTokenService"
	"app/repository"
	"app/service/noExpireTokenProvider"
	"app/valueObject"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	exp_duration   = time.Minute * 5
	claim_subject  = "sub"
	claim_audience = "aud"
	claim_expire   = "exp"

	//key_no_expire = "NO_EXPIRE_TOKEN"
)

var (
	ERR_INVALID_ACCESS_TOKEN_TYPE = errors.New("invalid access token format")
	ERR_INVALID_USER_CREDENTIALS  = errors.New("accessToken error: invalid username or password")
)

type (
	IAccessTokenHandler = accessTokenServicePort.IAccessTokenManipulator
	IGeneralToken       = generalTokenServicePort.IGeneralToken

	JWTAccessTokenManipulatorService struct {
		AudienceList               accessTokenServicePort.AccessTokenAudienceList
		JWTTokenManipulatorService jwtTokenServicePort.IAsymmetricJWTTokenManipulator
		GetUserAuthority           authServicePort.IGetUserAuthorityServicePort
		UserRepo                   repository.IUser
		ExpDuration                time.Duration
		WithoutExpire              bool
		noExpireTokenProvider.NoExpireTokenProvider
	}
)

func New(options ...AccessTokenManipulatorOption) *JWTAccessTokenManipulatorService {

	ret := new(JWTAccessTokenManipulatorService)

	for _, fn := range options {

		fn(ret)
	}

	return ret
}

func (this *JWTAccessTokenManipulatorService) Read(token_str string) (IAccessToken, error) {

	token, err := this.JWTTokenManipulatorService.VerifyTokenStringCustomClaim(token_str, &jwt_access_token_custom_claims{})

	if err != nil {

		return nil, err
	}

	return newFromToken(token)
}

func (this *JWTAccessTokenManipulatorService) GenerateBased(
	accessToken IAccessToken, ctx context.Context,
) (IAccessToken, error) {

	newAt, err := this.makeFor(accessToken.GetTenantUUID(), ctx)

	if err != nil {

		return nil, err
	}

	if v, ok := (accessToken.GetAuthData()).(*valueObject.AuthData); ok {

		newAt.customClaims.AuthData = v

	} else {

		return nil, libError.NewInternal(ERR_INVALID_TOKEN)
	}

	return newAt, nil
}

func (this *JWTAccessTokenManipulatorService) makeFor(
	tenantUUID uuid.UUID, ctx context.Context,
) (*jwt_access_token, error) {

	token := this.JWTTokenManipulatorService.GenerateToken()

	customeClaims := &jwt_access_token_custom_claims{
		Issuer:     bootstrap.GetAppName(),
		IssuedAt:   jwt.NewNumericDate(time.Now()),
		TokenID:    "",
		TenantUUID: libCommon.PointerPrimitive(tenantUUID),
		AuthData:   nil,
	}

	if !this.WithoutExpire && !this.IsNoExpire(ctx) {

		switch {
		case this.ExpDuration > 0:
			customeClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(this.ExpDuration))
		case this.ExpDuration <= 0:
			customeClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(exp_duration))
		}
	}

	token.Claims = customeClaims

	accesstoken, err := newFromToken(token)

	if err != nil {

		return nil, err
	}

	return accesstoken, nil
}

func (this *JWTAccessTokenManipulatorService) SignString(accessToken IAccessToken) (string, error) {

	if val, ok := accessToken.(*jwt_access_token); ok {

		return this.JWTTokenManipulatorService.SignString(val.jwt_token)
	}

	return "", libError.NewInternal(ERR_INVALID_ACCESS_TOKEN_TYPE)
}

func (this *JWTAccessTokenManipulatorService) DefaultExpireDuration() time.Duration {

	return exp_duration
}

func (this *JWTAccessTokenManipulatorService) GenerateFor(
	tenantUUID uuid.UUID, generalToken IGeneralToken, tokenID string, ctx context.Context,
) (accessTokenServicePort.IAccessToken, error) {

	authData, err := this.GetUserAuthority.Serve(tenantUUID, generalToken.GetUserUUID(), ctx)

	if err != nil {

		return nil, err
	}

	if authData == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	at, err := this.makeFor(tenantUUID, ctx)

	if err != nil {

		return nil, err
	}

	at.SetTokenID(tokenID)
	at.customClaims.AuthData = authData

	fmt.Println(authData.TenantAgentData)
	return at, nil
}
