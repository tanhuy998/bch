package refreshTokenClientPort

import (
	refreshTokenServicePort "app/port/refreshToken"
	"context"
)

type (
	IRefreshTokenClient interface {
		Read(ctx context.Context) (refreshTokenServicePort.IRefreshToken, error)
		Write(ctx context.Context, refreshToken refreshTokenServicePort.IRefreshToken) error
	}
)
