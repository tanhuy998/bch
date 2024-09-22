package accessTokenClientPort

import "github.com/kataras/iris/v12"

type (
	IAccessTokenClient interface {
		Read(reqCtx iris.Context) string
	}
)
