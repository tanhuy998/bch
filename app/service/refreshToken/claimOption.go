package refreshTokenService

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SetExpire(t time.Time) ClaimsOption {

	return func(claim *jwt.RegisteredClaims) {

		claim.ExpiresAt = jwt.NewNumericDate(t)
	}
}
