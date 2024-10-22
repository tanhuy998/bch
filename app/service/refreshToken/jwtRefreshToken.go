package refreshTokenService

import (
	libError "app/internal/lib/error"
	refreshTokenIdServicePort "app/port/refreshTokenID"
	"errors"
	"fmt"
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
		Issuer         string           `json:"iss"`
		RefreshTokenID string           `json:"jti"`
		TenantUUID     *uuid.UUID       `json:"sub"`
		IssuedAt       *jwt.NumericDate `json:"iat"`
		ExpiresAt      *jwt.NumericDate `json:"exp"`
	}

	jwt_refresh_token struct {
		userUUID  uuid.UUID
		jwt_token *jwt.Token
		claims    *jwt_refresh_token_custom_claims
	}
)

func newFromToken(
	token *jwt.Token, refreshTokenIDProvider refreshTokenIdServicePort.IRefreshTokenIDProvider,
) (*jwt_refresh_token, error) {

	var (
		ret *jwt_refresh_token
	)

	if claims, ok := token.Claims.(*jwt_refresh_token_custom_claims); ok {

		ret = &jwt_refresh_token{
			jwt_token: token,
			claims:    claims,
		}

	} else {

		return nil, ERR_INVALID_TOKEN
	}

	claims := ret.claims

	if claims.TenantUUID == nil ||
		*claims.TenantUUID == uuid.Nil {

		return nil, libError.NewInternal(fmt.Errorf("refresh token contains no Tenant UUID"))
	}

	if claims.RefreshTokenID == "" {

		return nil, libError.NewInternal(fmt.Errorf("refresh token contains no token ID"))
	}

	generalTokenID, _, err := refreshTokenIDProvider.Extract(claims.RefreshTokenID)

	if err != nil {

		return nil, err
	}

	if generalTokenID.GetUserUUID() == uuid.Nil {

		return nil, libError.NewInternal(fmt.Errorf("refresh token contain invalid token id"))
	}

	ret.userUUID = generalTokenID.GetUserUUID()

	return ret, nil
}

func (this *jwt_refresh_token) GetUserUUID() uuid.UUID {

	return this.userUUID
}
func (this *jwt_refresh_token) GetTokenID() string {

	return this.claims.RefreshTokenID
}

func (this *jwt_refresh_token) Expired() bool {

	exp, err := this.jwt_token.Claims.GetExpirationTime()

	if err != nil {

		return false
	}

	if exp == nil {

		return false
	}

	return exp.Before(time.Now())
}

func (this *jwt_refresh_token) GetExpireTime() *time.Time {

	claims := this.claims

	if claims == nil ||
		claims.ExpiresAt == nil {

		return nil
	}

	return &claims.ExpiresAt.Time
}

func (this *jwt_refresh_token) GetTenantUUID() uuid.UUID {

	return *this.claims.TenantUUID
}
