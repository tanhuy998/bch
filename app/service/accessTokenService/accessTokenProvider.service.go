package accessTokenService

import (
	accessTokenServicePort "app/adapter/accessToken"
	authServiceAdapter "app/adapter/auth"
	jwtTokenServicePort "app/adapter/jwtTokenService"
	"app/repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	exp_duration   = time.Minute * 5
	claim_subject  = "sub"
	claim_audience = "aud"
	claim_expire   = "exp"
)

var (
	ERR_INVALID_ACCESS_TOKEN_TYPE = errors.New("invalid access token format")
)

type (
	IAccessTokenHandler = accessTokenServicePort.IAccessTokenManipulator

	JWTAccessTokenManipulatorService struct {
		//JWtTokenSigningServive      jwtTokenServicePort.IJWTTokenSigning
		AudienceList               accessTokenServicePort.AccessTokenAudienceList
		FetchAuthDataService       authServiceAdapter.IFetchAuthData
		JWTTokenManipulatorService jwtTokenServicePort.IAsymmetricJWTTokenManipulator
		UserRepo                   repository.IUser
	}
)

func (this *JWTAccessTokenManipulatorService) Read(token_str string) (IAccessToken, error) {

	token, err := this.JWTTokenManipulatorService.VerifyTokenString(token_str)

	if err != nil {

		return nil, err
	}

	return newFromToken(token)
}

func (this *JWTAccessTokenManipulatorService) Generate(userUUID uuid.UUID, ctx context.Context) (accessTokenServicePort.IAccessToken, error) {

	authData, err := this.FetchAuthDataService.Serve(userUUID, ctx)

	if err != nil {

		return nil, err
	}

	accessToken, err := this.makeFor(userUUID)

	if err != nil {

		return nil, err
	}

	accessToken.claims.AuthData = authData

	return accessToken, nil
}

func (this *JWTAccessTokenManipulatorService) makeFor(userUUID uuid.UUID) (*jwt_access_token, error) {

	token := this.JWTTokenManipulatorService.GenerateToken()

	token.Claims = jwt_access_token_custom_claims{
		jwt.RegisteredClaims{
			Subject:   userUUID.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp_duration)),
			Audience:  jwt.ClaimStrings(this.AudienceList),
			//Issuer: ,
		},
		nil,
	}

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

	return "", ERR_INVALID_ACCESS_TOKEN_TYPE
}
