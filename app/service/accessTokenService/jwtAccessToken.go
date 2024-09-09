package accessTokenService

import (
	accessTokenServicePort "app/adapter/accessToken"
	libCommon "app/lib/common"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ERR_INVALID_TOKEN = errors.New("invalid jwt token")
)

type (
	IAccessToken interface {
		accessTokenServicePort.IAccessToken
	}

	jwt_access_token_custom_claims struct {
		// Sub *uuid.UUID `json:"sub"`
		// Aud []string   `json:"aud,omitempty"`
		// Exp time.Time  `json:"exp"`
		jwt.RegisteredClaims
	}

	jwt_access_token struct {
		jwt_token *jwt.Token
		claims    *jwt_access_token_custom_claims
		userUUID  *uuid.UUID
		expired   bool
		// audience  *accessTokenServicePort.AccessTokenAudience
	}
)

func NewFromToken(token *jwt.Token) (*jwt_access_token, error) {

	var ret *jwt_access_token

	if val, ok := token.Claims.(*jwt_access_token_custom_claims); !ok {

		ret = &jwt_access_token{
			jwt_token: token,
			claims:    val,
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

	ret.expired = exp.After(time.Now())

	return ret, nil
}

func (this *jwt_access_token) GetUserUUID() *uuid.UUID {

	return this.userUUID
}

func (this *jwt_access_token) GetAudiences() []string {

	return nil
}

func (this *jwt_access_token) Expired() bool {

	return this.expired
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
