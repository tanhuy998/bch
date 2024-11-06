package authSignaturesServicePort

import (
	accessTokenServicePort "app/port/accessToken"
	generalTokenServicePort "app/port/generalToken"

	refreshTokenServicePort "app/port/refreshToken"
	"context"

	"github.com/google/uuid"
)

type (
	ISwitchTenant interface {
		Serve(
			tenantUUID uuid.UUID, generalToken generalTokenServicePort.IGeneralToken, ctx context.Context,
		) (at accessTokenServicePort.IAccessToken, rt refreshTokenServicePort.IRefreshToken, err error)
	}
)
