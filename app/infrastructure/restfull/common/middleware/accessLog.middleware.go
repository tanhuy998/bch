package middleware

import (
	"app/internal/bootstrap"
	"app/internal/errorMap"
	accessLogServicePort "app/port/accessLog"
	"app/port/envConsumerServicePort"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

type (
	InternalAccessLogHandler struct {
		AccessLogger accessLogServicePort.IAccessLogger
		ErrorHandler IErrorHandler
		Env          envConsumerServicePort.IENVConsumer
	}
)

func (this *InternalAccessLogHandler) handlePanic(ctx iris.Context, r interface{}) {

	switch actualRecoverVal := r.(type) {
	case error:
		this.pushRecoverError(ctx, actualRecoverVal)
	case string:
		this.wrapPanicMessage(ctx, actualRecoverVal)
	case fmt.Stringer:
		this.wrapPanicMessage(ctx, actualRecoverVal.String())
	default:
		this.pushRecoverError(ctx, fmt.Errorf("%v", actualRecoverVal))
	}
}

func (this *InternalAccessLogHandler) pushRecoverError(ctx iris.Context, err error) {

	switch {
	case !errors.Is(err, errorMap.ERR_INTERNAL):
		err = errorMap.WrapAndExportAsInternalError(err)
		fallthrough
	default:
		this.AccessLogger.PushError(ctx, err)

		if isDebugging, ok := this.Env.Get(bootstrap.ENV_DEBUG_LOG).ToBoolean(); ok && isDebugging {

			this.AccessLogger.PushError(
				ctx, fmt.Errorf("%s", debug.Stack()),
			)
		}

		this.ErrorHandler.HandleContextError(ctx, err)
	}
}

func (this *InternalAccessLogHandler) wrapPanicMessage(ctx iris.Context, msg string) {

	this.pushRecoverError(
		ctx,
		errorMap.WrapAndExportAsInternalError(
			fmt.Errorf(msg),
		),
	)
}

func InternalAccessLog(container *hero.Container) context.Handler {

	return container.Handler(
		func(ctx iris.Context, handler *InternalAccessLogHandler) {

			defer func() {

				if r := recover(); r != nil {

					handler.handlePanic(ctx, r)
				}

				handler.AccessLogger.EndContext(ctx)
			}()

			handler.AccessLogger.Init(ctx)

			ctx.Next()
		},
	)
}
