package accessTokenService

import (
	accessTokenServicePort "app/adapter/accessToken"
	"app/domain/valueObject"
	libCommon "app/lib/common"
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
		TokenID  string                `json:"jti"`
		AuthData *valueObject.AuthData `json:"aut"`
	}

	jwt_access_token struct {
		jwt_token    *jwt.Token
		customClaims *jwt_access_token_custom_claims
		userUUID     *uuid.UUID
		//authData  *AccessTokenAuthData
		//expired bool
		// audience  *accessTokenServicePort.AccessTokenAudience
	}
)

func newFromToken(token *jwt.Token) (*jwt_access_token, error) {

	if token == nil {

		return nil, ERR_NIL_TOKEN
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

	subClaim, err := token.Claims.GetSubject()

	if err != nil {

		return nil, err
	}

	userUUID, err := uuid.Parse(subClaim)

	if err != nil {

		return nil, err
	}

	ret.userUUID = libCommon.PointerPrimitive(userUUID)

	exp, err := token.Claims.GetExpirationTime()

	if err != nil {

		return nil, err
	}

	if exp == nil {

		return nil, ERR_INVALID_TOKEN
	}

	return ret, nil
}

func (this *jwt_access_token) GetUserUUID() uuid.UUID {

	return *this.userUUID
}

func (this *jwt_access_token) GetAudiences() []string {

	return nil
}

func (this *jwt_access_token) Expired() bool {

	exp, err := this.GetExpire()

	if err != nil {

		return false
	}

	return exp.Before(time.Now())
}

func (this *jwt_access_token) GetAuthData() IAccessTokenAuthData {

	return this.customClaims.AuthData
}

func (this *jwt_access_token) GetExpire() (*jwt.NumericDate, error) {

	return this.jwt_token.Claims.GetExpirationTime()
}

func (this *jwt_access_token) GetTokenID() string {

	return this.customClaims.TokenID
}

func (this *jwt_access_token) SetTokenID(id string) {

	this.customClaims.TokenID = id
}

// func (this *jwt_access_token) GetParsedAudience() *accessTokenServicePort.AccessTokenAudience {

// 	return this.audience
// }

// func (this *jwt_access_token) GetExpirationTime() time.Time {

// 	t, _ := this.jwt_token.Claims.GetExpirationTime()

// 	return t.Time
// }

// func (this *jwt_access_token) GetIssuedAt() time.Time {

// 	i, _ := this.jwt_token.Claims.GetIssuedAt()

// 	return i.Time
// }

// func (this *jwt_access_token) GetNotBefore() time.Time {

// 	v, _ := this.jwt_token.Claims.GetNotBefore()

// 	return v.Time
// }

// func (this *jwt_access_token) GetIssuer() string {

// 	v, _ := this.jwt_token.Claims.GetIssuer()

// 	return v
// }

// func (this *jwt_access_token) GetSubject() *uuid.UUID {

// 	//v, _ := this.jwt_token.Claims.GetSubject()

// 	return this.userUUID
// }

// func (this *jwt_access_token) GetAudience() *accessTokenServicePort.AccessTokenAudience {

// 	return this.audience
// }
