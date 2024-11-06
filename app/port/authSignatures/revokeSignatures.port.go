package authSignaturesServicePort

import (
	accessTokenServicePort "app/port/accessToken"
	refreshTokenServicePort "app/port/refreshToken"
	"context"
)

type (
	IRevokeSignatures interface {
		Serve(
			refreshToken refreshTokenServicePort.IRefreshToken, accessToken accessTokenServicePort.IAccessToken, ctx context.Context,
		) error
	}
)
