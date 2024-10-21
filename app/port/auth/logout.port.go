package authServicePort

import (
	accessTokenServicePort "app/port/accessToken"
	refreshTokenServicePort "app/port/refreshToken"
	"context"
)

type (
	ILogout interface {
		Serve(
			refreshToken refreshTokenServicePort.IRefreshToken, accessToken accessTokenServicePort.IAccessToken, ctx context.Context,
		) error
	}
)
