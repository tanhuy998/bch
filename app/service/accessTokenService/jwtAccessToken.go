package accessTokenService

import (
	libError "app/internal/lib/error"
	accessTokenServicePort "app/port/accessToken"
	"app/valueObject"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ERR_INVALID_TOKEN = errors.New("invalid jwt token")
	ERR_NIL_TOKEN     = errors.New("nil token")
)

type (
	IAccessToken         = accessTokenServicePort.IAccessToken
	IAccessTokenAuthData = accessTokenServicePort.IAccessTokenAuthData

	// IAccessTokenAuthData interface {
	// 	accessTokenServicePort.IAccessTokenAuthData
	// }

	jwt_access_token_custom_claims struct {
		// Sub *uuid.UUID `json:"sub"`
		// Aud []string   `json:"aud,omitempty"`
		// Exp time.Time  `json:"exp"`
		jwt.RegisteredClaims
		Issuer     string                `json:"iss"`
		TokenID    string                `json:"jti"`
		IssuedAt   *jwt.NumericDate      `json:"iat"`
		ExpireAt   *jwt.NumericDate      `json:"exp"`
		TenantUUID *uuid.UUID            `json:"sub"`
		AuthData   *valueObject.AuthData `json:"aut"`
	}

	jwt_access_token struct {
		jwt_token    *jwt.Token
		customClaims *jwt_access_token_custom_claims
		// userUUID     *uuid.UUID
		//authData  *AccessTokenAuthData
		//expired bool
		// audience  *accessTokenServicePort.AccessTokenAudience
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

	// subClaim, err := token.Claims.GetSubject()

	// if err != nil {

	// 	return nil, errors.Join(
	// 		common.ERR_INTERNAL,
	// 		err,
	// 	)
	// }

	// userUUID, err := uuid.Parse(subClaim)

	// if err != nil {

	// 	return nil, errors.Join(
	// 		common.ERR_INTERNAL,
	// 		err,
	// 	)
	// }

	// ret.userUUID = libCommon.PointerPrimitive(userUUID)

	// exp, err := token.Claims.GetExpirationTime()

	// if err != nil {

	// 	return nil, err
	// }

	// if exp == nil {

	// 	return nil, ERR_INVALID_TOKEN
	// }

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
