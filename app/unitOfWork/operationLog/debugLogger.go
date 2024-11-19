package opLog

import (
	"app/internal/bootstrap"
	accessLogServicePort "app/port/accessLog"
	"context"
	"os"
)

type (
	DebugLogger struct {
		AccessLogger accessLogServicePort.IAccessLogger
	}
)

func (this *DebugLogger) couldLog(ctx context.Context) bool {

	return os.Getenv(bootstrap.ENV_DEBUG_LOG) == "true" && couldLog(this.AccessLogger, ctx)
}

func (this *DebugLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	if !this.couldLog(ctx) {

		return empty_trace_func
	}

	return messure(this.AccessLogger, LOG_LEVEL_DEBUG, op, msg, ctx)
}

func (this *DebugLogger) PushIfError(err error, op string, msg string, ctx context.Context) {

	if err == nil || !this.couldLog(ctx) {

		return
	}

	pushTraceIfError(this.AccessLogger, LOG_LEVEL_DEBUG, err, op, msg, ctx)
}

func (this *DebugLogger) Push(op string, msg string, ctx context.Context) {

	if !this.couldLog(ctx) {

		return
	}

	pushTrace(this.AccessLogger, LOG_LEVEL_DEBUG, op, msg, ctx)
}

func (this *DebugLogger) PushCond(
	op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	if !this.couldLog(ctx) {

		return empty_push_cond_func
	}

	return pushTraceCond(this.AccessLogger, LOG_LEVEL_DEBUG, op, msgIfNoErr, ctx)
}

func (this *DebugLogger) PushCondWithMessurement(
	op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	if !this.couldLog(ctx) {

		return empty_push_cond_with_messurement_func
	}

	return PushTraceCondWithMessurement(
		this.AccessLogger, LOG_LEVEL_DEBUG, op, ctx,
	)
}

func (this *DebugLogger) PushError(op string, err error, defaultMsg string, ctx context.Context) {

	if !this.couldLog(ctx) {

		return
	}

	pushTraceError(this.AccessLogger, LOG_LEVEL_DEBUG, op, err, defaultMsg, ctx)
}
