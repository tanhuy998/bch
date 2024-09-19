package refreshTokenIdServicePort

import "github.com/google/uuid"

const (
	REFRESH_TOKEN_COOKIE          = "refresh-token"
	REFRESH_TOKEN_COOKIE_DURATION = 2
)

type (
	IRefreshTokenIDProvider interface {
		Generate(userUUID uuid.UUID) (string, error)
	}
)
