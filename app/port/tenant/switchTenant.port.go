package tenantServicePort

import (
	accessTokenServicePort "app/port/accessToken"
	"app/port/generalTokenServicePort"
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
