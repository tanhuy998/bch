package authSignatureTokenPort

import (
	accessTokenServicePort "app/src/port/accessToken"
	refreshTokenServicePort "app/src/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	IAuthSignatureProvider interface {
		Generate(
			userUUID uuid.UUID, ctx context.Context,
		) (accessTokenServicePort.IAccessToken, refreshTokenServicePort.IRefreshToken, error)
		Rotate(
			refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context,
		) (accessTokenServicePort.IAccessToken, refreshTokenServicePort.IRefreshToken, error)
		GenerateStrings(
			userUUID uuid.UUID, ctx context.Context,
		) (at string, rt string, err error)
		RotateStrings(
			refreshToken refreshTokenServicePort.IRefreshToken, ctx context.Context,
		) (at string, rt string, err error)
	}
)
