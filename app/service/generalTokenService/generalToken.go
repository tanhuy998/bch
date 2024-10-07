package generalTokenService

import (
	libError "app/internal/lib/error"
	tenantAccessTokenServicePort "app/port/generalTokenServicePort"
	"fmt"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	ITenantAccessToken = tenantAccessTokenServicePort.IGeneralToken

	custom_claims struct {
		jwt.RegisteredClaims
		UserUUID *uuid.UUID `json:"sub"`
	}

	jwt_tenant_access_token struct {
		jwt_token *jwt.Token
		claims    *custom_claims
		// userUUID   *uuid.UUID
		// tenantUUID *uuid.UUID
	}
)

func newFromToken(token *jwt.Token) (*jwt_tenant_access_token, error) {

	if token == nil {

		return nil, libError.NewInternal(fmt.Errorf("tenant access token error: nil token given"))
	}

	custom_claims, ok := token.Claims.(*custom_claims)

	if !ok {

		return nil, libError.NewInternal(fmt.Errorf("tenant access token err: invalid token claim type"))
	}

	ret := &jwt_tenant_access_token{
		jwt_token: token,
		claims:    custom_claims,
	}

	return ret, nil
}

func (this *jwt_tenant_access_token) GetUserUUID() uuid.UUID {

	return *this.claims.UserUUID
}

func (this *jwt_tenant_access_token) GetExpiretime() *time.Time {

	t, err := this.claims.GetExpirationTime()

	if err != nil {

		return nil
	}

	return &t.Time
}
