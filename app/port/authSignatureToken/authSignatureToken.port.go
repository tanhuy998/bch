package authSignatureTokenPort

import (
	accessTokenServicePort "app/port/accessToken"
	"app/port/generalTokenServicePort"
	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	IGeneralToken = generalTokenServicePort.IGeneralToken

	IAuthSignatureProvider interface {
		Generate(
			tenantUUID uuid.UUID, generalToken IGeneralToken, ctx context.Context,
		) (accessTokenServicePort.IAccessToken, refreshTokenServicePort.IRefreshToken, error)
		Rotate(
			oldRefreshToken refreshTokenServicePort.IRefreshToken, oldAccessToken accessTokenServicePort.IAccessToken, ctx context.Context,
		) (accessTokenServicePort.IAccessToken, refreshTokenServicePort.IRefreshToken, error)
		GenerateStrings(
			tenantUUID uuid.UUID, generalToken IGeneralToken, ctx context.Context,
		) (at string, rt string, err error)
		RotateStrings(
			oldRefreshToken refreshTokenServicePort.IRefreshToken, oldAccessToken accessTokenServicePort.IAccessToken, ctx context.Context,
		) (at string, rt string, err error)
	}
)
