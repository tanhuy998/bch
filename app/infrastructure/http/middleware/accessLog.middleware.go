package middleware

import (
	"app/infrastructure/http/api/v1/config"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

func InternalAccessLog(container *hero.Container) context.Handler {

	return container.Handler(
		func(ctx iris.Context, accessLogger config.AccessLogger) {

			accessLogger.Init(ctx)

			ctx.Next()

			go accessLogger.EndContext(ctx)
		},
	)
}
