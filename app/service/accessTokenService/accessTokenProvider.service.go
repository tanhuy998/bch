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
	jwtClaim "app/valueObject/jwt"
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	default_exp_duration = time.Minute * 1
	claim_subject        = "sub"
	claim_audience       = "aud"
	claim_expire         = "exp"

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

	newAt, err := this.makeFor(accessToken.GetTenantUUID(), accessToken.GetUserUUID(), ctx)

	if err != nil {

		return nil, err
	}

	return newAt, nil
}

func (this *JWTAccessTokenManipulatorService) makeFor(
	tenantUUID, userUUID uuid.UUID, ctx context.Context,
) (*jwt_access_token, error) {

	authData, err := this.GetUserAuthority.Serve(tenantUUID, userUUID, ctx)

	if err != nil {

		return nil, err
	}

	if authData == nil {

		return nil, common.ERR_UNAUTHORIZED
	}

	token := this.JWTTokenManipulatorService.GenerateToken()

	customClaims := &jwt_access_token_custom_claims{
		Issuer:     bootstrap.GetAppName(),
		IssuedAt:   jwt.NewNumericDate(time.Now()),
		TokenID:    "",
		TenantUUID: libCommon.PointerPrimitive(tenantUUID),
		AuthData:   authData,
		ExpireAt:   jwt.NewNumericDate(time.Now().Add(default_exp_duration)),
	}

	jwtClaim.SetupAccessToken(&customClaims.PrivateClaims)

	if this.WithoutExpire || this.IsNoExpire(ctx) {

		customClaims.ExpireAt = nil
	}

	if os.Getenv("TEST-LOGIN") == "1" {

		customClaims.ExpireAt = jwt.NewNumericDate(time.Now())
	}

	token.Claims = customClaims

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

	return default_exp_duration
}

func (this *JWTAccessTokenManipulatorService) GenerateFor(
	tenantUUID uuid.UUID, generalToken IGeneralToken, tokenID string, ctx context.Context,
) (accessTokenServicePort.IAccessToken, error) {

	at, err := this.makeFor(tenantUUID, generalToken.GetUserUUID(), ctx)

	if err != nil {

		return nil, err
	}

	fmt.Println(at.customClaims.AuthData)

	switch policies := generalToken.GetPolicies(); {
	case policies == nil:
	case slices.Contains(policies, jwtClaim.POLICY_AT_NO_EXPIRE):
		at.customClaims.ExpireAt = nil
	}

	at.SetTokenID(tokenID)
	// at.customClaims.AuthData = authData

	// fmt.Println(authData.TenantAgentData)
	return at, nil
}
