package accessTokenService

import (
	accessTokenServicePort "app/adapter/accessToken"
	jwtTokenServicePort "app/adapter/jwtTokenService"
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
	IAccessTokenHandler interface {
		accessTokenServicePort.IAccessTokenHandler
	}

	JWTAccessTokenHandlerService struct {
		//JWtTokenSigningServive      jwtTokenServicePort.IJWTTokenSigning
		AudienceList               accessTokenServicePort.AccessTokenAudienceList
		JWTTokenManipulatorService jwtTokenServicePort.IJWTTokenManipulator
	}
)

func (this *JWTAccessTokenHandlerService) Read(token_str string) (IAccessToken, error) {

	token, err := this.JWTTokenManipulatorService.VerifyTokenString(token_str)

	if err != nil {

		return nil, err
	}

	return NewFromToken(token)
}

func (this *JWTAccessTokenHandlerService) Generate(subject uuid.UUID) (accessTokenServicePort.IAccessToken, error) {

	token := this.JWTTokenManipulatorService.GenerateToken()

	token.Claims = jwt_access_token_custom_claims{
		jwt.RegisteredClaims{
			Subject:   subject.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp_duration)),
			Audience:  jwt.ClaimStrings(this.AudienceList),
		},
	}

	accesstoken, err := NewFromToken(token)

	if err != nil {

		return nil, err
	}

	return accesstoken, nil
}

func (this *JWTAccessTokenHandlerService) SignedString(accessToken IAccessToken) (string, error) {

	if val, ok := accessToken.(*jwt_access_token); ok {

		return this.JWTTokenManipulatorService.SignedString(val.jwt_token)
	}

	return "", ERR_INVALID_ACCESS_TOKEN_TYPE
}
