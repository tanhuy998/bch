package accessTokenClientPort

import (
	accessTokenServicePort "app/port/accessToken"

	"github.com/kataras/iris/v12"
)

type (
	IAccessTokenClient interface {
		Read(reqCtx iris.Context) (accessTokenServicePort.IAccessToken, error)
	}
)
