package refreshTokenService

import (
	libCommon "app/lib/common"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ERR_INVALID_TOKEN = errors.New("refeshToken error: invalid token")
)

type (
	jwt_refresh_token_custom_claims struct {
		jwt.RegisteredClaims
		RefreshTokenID string `json:"jti"`
	}

	jwt_refresh_token struct {
		jwt_token *jwt.Token
		claims    *jwt_refresh_token_custom_claims
		userUUID  *uuid.UUID
		expired   bool
	}
)

func newFromToken(token *jwt.Token) (*jwt_refresh_token, error) {

	var ret *jwt_refresh_token

	if claims, ok := token.Claims.(*jwt_refresh_token_custom_claims); ok {

		ret = &jwt_refresh_token{
			jwt_token: token,
			claims:    claims,
		}

	} else {

		return nil, ERR_INVALID_TOKEN
	}

	exp, err := token.Claims.GetExpirationTime()

	if err != nil {

		return nil, err
	}

	ret.expired = time.Now().After(exp.Time)

	subject, err := token.Claims.GetSubject()

	if err != nil {

		return nil, err
	}

	userUUID, err := uuid.Parse(subject)

	if err != nil {

		return nil, err
	}

	ret.userUUID = libCommon.PointerPrimitive(userUUID)

	uuid.New().ID()

	return ret, nil
}

func (this *jwt_refresh_token) GetUserUUID() uuid.UUID {

	return *this.userUUID
}
func (this *jwt_refresh_token) GetTokenID() string {

	return this.claims.RefreshTokenID
}
