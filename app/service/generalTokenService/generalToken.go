package generalTokenService

import (
	"app/internal/generalToken"
	libError "app/internal/lib/error"
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	custom_claims struct {
		jwt.RegisteredClaims
		IssuedAt       *jwt.NumericDate             `json:"iat"`
		ExpireAt       *jwt.NumericDate             `json:"exp"`
		GeneralTokenID *generalToken.GeneralTokenID `json:"jti"`
		UserUUID       *uuid.UUID                   `json:"sub"`
	}

	jwt_general_token struct {
		jwt_token *jwt.Token
		claims    *custom_claims
		// userUUID   *uuid.UUID
		// tenantUUID *uuid.UUID
	}
)

func newFromToken(token *jwt.Token) (*jwt_general_token, error) {

	if token == nil {

		return nil, libError.NewInternal(fmt.Errorf("general token error: nil token given"))
	}

	custom_claims, ok := token.Claims.(*custom_claims)

	if !ok {

		return nil, libError.NewInternal(fmt.Errorf("general token error: invalid token claim type"))
	}

	ret := &jwt_general_token{
		jwt_token: token,
		claims:    custom_claims,
	}

	if custom_claims.UserUUID == nil ||
		*custom_claims.UserUUID == uuid.Nil {

		return nil, libError.NewInternal(fmt.Errorf("general token contains no user uuid"))
	}

	if custom_claims.GeneralTokenID == nil ||
		*custom_claims.GeneralTokenID == generalToken.Nil {

		return nil, libError.NewInternal(fmt.Errorf("general token contains no token id"))
	}

	return ret, nil
}

func (this *jwt_general_token) GetUserUUID() uuid.UUID {

	return *this.claims.UserUUID
}

func (this *jwt_general_token) GetExpiretime() *time.Time {

	// t, err := this.claims.GetExpirationTime()

	// if err != nil {

	// 	return nil
	// }

	// if t == nil {

	// 	return nil
	// }

	// return &t.Time

	claims := this.claims

	if claims == nil ||
		claims.ExpiresAt == nil {

		return nil
	}

	return &claims.ExpiresAt.Time
}

func (this *jwt_general_token) GetTokenID() generalToken.GeneralTokenID {

	return *this.claims.GeneralTokenID
}

func (this *jwt_general_token) Expire() bool {

	exp := this.GetExpiretime()

	if exp == nil {

		return false
	}

	return time.Now().After(*exp)
}
