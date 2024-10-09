package refreshTokenIdServicePort

import (
	"app/internal/generalToken"
)

const (
	REFRESH_TOKEN_COOKIE          = "refresh-token"
	REFRESH_TOKEN_COOKIE_DURATION = 2
)

type (
	IRefreshTokenIDProvider interface {
		Generate(generalTokenID generalToken.GeneralTokenID) (string, error)
		Extract(
			ID string,
		) (generalTokenID generalToken.GeneralTokenID, sessionID [8]byte, err error)
	}
)
