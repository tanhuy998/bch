package middleware

import (
	accessLogServicePort "app/port/accessLog"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

func InternalAccessLog(container *hero.Container) context.Handler {

	return container.Handler(
		func(ctx iris.Context, accessLogger accessLogServicePort.IAccessLogger) {

			defer func() {

				if r := recover(); r != nil {

					defer panic(r)
				}

				accessLogger.EndContext(ctx)
			}()

			accessLogger.Init(ctx)

			ctx.Next()
		},
	)
}
