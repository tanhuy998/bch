package refreshTokenClientPort

import "github.com/kataras/iris/v12"

type (
	IRefreshTokenClient interface {
		Read(ctx iris.Context) string
		Write(ctx iris.Context, refreshToken string) error
	}
)
