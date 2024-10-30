package accessTokenService

import (
	"app/internal/common"
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	"app/valueObject"
	jwtClaim "app/valueObject/jwt"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ERR_INVALID_TOKEN = fmt.Errorf("invalid jwt token")
	ERR_NIL_TOKEN     = fmt.Errorf("nil token")
)

type (
	IAccessToken         = accessTokenServicePort.IAccessToken
	IAccessTokenAuthData = accessTokenServicePort.IAccessTokenAuthData

	jwt_access_token_custom_claims struct {
		jwt.RegisteredClaims
		Issuer  string `json:"iss,omitempty"`
		TokenID string `json:"jti,omitempty"`
		jwtClaim.PrivateClaims
		IssuedAt   *jwt.NumericDate      `json:"iat,omitempty"`
		ExpireAt   *jwt.NumericDate      `json:"exp,omitempty"`
		TenantUUID *uuid.UUID            `json:"sub,omitempty"`
		AuthData   *valueObject.AuthData `json:"aut,omitempty"`
	}

	jwt_access_token struct {
		jwt_token    *jwt.Token
		customClaims *jwt_access_token_custom_claims
	}
)

func newFromToken(token *jwt.Token) (*jwt_access_token, error) {

	if token == nil {

		return nil, libError.NewInternal(ERR_NIL_TOKEN)
	}

	var ret *jwt_access_token

	if val, ok := token.Claims.(*jwt_access_token_custom_claims); ok {

		ret = &jwt_access_token{
			jwt_token:    token,
			customClaims: val,
		}

	} else {

		return nil, ERR_INVALID_TOKEN
	}

	claims := ret.customClaims

	switch {
	case claims.TokenType != jwtClaim.ACCESS_TOKEN:
		return nil, errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("the given token is not access token"),
		)
	case claims.TenantUUID == nil || *claims.TenantUUID == uuid.Nil:
		return nil, errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("the given token has no subject"),
		)
	case claims.AuthData == nil:
		return nil, errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("the given token has no authority"),
		)
	case ret.GetUserUUID() == uuid.Nil:
		return nil, errors.Join(
			common.ERR_UNAUTHORIZED, fmt.Errorf("the given token has no user"),
		)
	}

	return ret, nil
}

func (this *jwt_access_token) GetUserUUID() uuid.UUID {

	claims := this.customClaims

	if claims == nil ||
		claims.AuthData == nil ||
		claims.AuthData.UserUUID == nil {

		return uuid.Nil
	}

	return *claims.AuthData.UserUUID
}

func (this *jwt_access_token) GetAudiences() []string {

	return nil
}

func (this *jwt_access_token) Expired() bool {

	exp := this.GetExpire()

	if exp == nil {

		return false
	}

	return exp.Before(time.Now())
}

func (this *jwt_access_token) GetAuthData() IAccessTokenAuthData {

	return this.customClaims.AuthData
}

func (this *jwt_access_token) GetExpire() *time.Time {

	claims := this.customClaims

	if claims == nil ||
		claims.ExpireAt == nil {

		return nil
	}

	return &claims.ExpireAt.Time
}

func (this *jwt_access_token) GetTokenID() string {

	return this.customClaims.TokenID
}

func (this *jwt_access_token) SetTokenID(id string) {

	this.customClaims.TokenID = id
}

func (this *jwt_access_token) GetTenantUUID() uuid.UUID {

	claims := this.customClaims

	if claims == nil ||
		claims.AuthData == nil ||
		claims.AuthData.TenantUUID == nil {

		return uuid.Nil
	}

	return *this.customClaims.TenantUUID
}
