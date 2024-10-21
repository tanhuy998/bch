package accessTokenClientPort

import (
	accessTokenServicePort "app/port/accessToken"
	"context"
)

type (
	IAccessTokenClient interface {
		Read(reqCtx context.Context) (accessTokenServicePort.IAccessToken, error)
	}
)
