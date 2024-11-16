package opLog

import (
	"app/internal/bootstrap"
	libCommon "app/internal/lib/common"
	accessLogServicePort "app/port/accessLog"
	"context"
	"os"
	"time"
)

const (
	ENV_OP_TRACE_DURATION = bootstrap.ENV_OP_TRACE_DURATION
)

type (
	IOperationLogger interface {
		Messure(op string, msg string, ctx context.Context) func(err error)
		PushTraceIfError(err error, op string, msg string, ctx context.Context)
		PushTrace(op string, msg string, ctx context.Context)
		PushTraceCond(op string, msgIfNoErr string, ctx context.Context) (logErrFunc func(err error, msgIfErr string))
		PushTraceCondWithMessurement(
			op string, msgIfNoErr string, ctx context.Context,
		) func(err error, msgIfErr string)
		PushTraceError(op string, err error, defaultMsg string, ctx context.Context)
	}

	OperationLogger struct {
		AccessLogger accessLogServicePort.IAccessLogger
	}
)

func (this *OperationLogger) CouldLog(ctx context.Context) bool {

	return ctx != nil && this.AccessLogger.IsLogging(ctx) && this.AccessLogger.IsTraceLogging(ctx)
}

func (this *OperationLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	if !this.CouldLog(ctx) {

		return empty_trace_func
	}

	start := time.Now()

	return func(err error) {

		duration := time.Since(start)

		l := &logPattern{
			Operation: op,
			Message:   msg,
		}

		l.SetDuration(duration)

		this.AccessLogger.PushTraceLogs(ctx, l)
	}
}

func (this *OperationLogger) PushTraceIfError(err error, op string, msg string, ctx context.Context) {

	if err == nil || !this.CouldLog(ctx) {

		return
	}

	this.AccessLogger.PushTraceLogs(
		ctx,
		logPattern{
			Operation: op,
			Message:   msg,
			ErrorMsg:  err.Error(),
		},
	)
}

func (this *OperationLogger) PushTrace(op string, msg string, ctx context.Context) {

	if !this.CouldLog(ctx) {

		return
	}

	this.AccessLogger.PushTraceLogs(
		ctx,
		logPattern{
			Operation: op,
			Message:   msg,
		},
	)
}

func (this *OperationLogger) PushTraceCond(
	op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	return func(err error, errMsg string) {

		if !this.CouldLog(ctx) {

			return
		}

		if err == nil {

			this.PushTrace(op, msgIfNoErr, ctx)
			return
		}

		this.PushTraceError(op, err, msgIfNoErr, ctx)
	}
}

func (this *OperationLogger) PushTraceCondWithMessurement(
	op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	start := time.Now()

	return func(msgIfNoErr string, err error, msgIfErr string) {

		if !this.CouldLog(ctx) {

			return
		}

		l := logPattern{
			Operation: op,
		}

		if err == nil {

			l.Message = msgIfNoErr

		} else {

			l.Message = libCommon.Ternary(msgIfErr == "", err.Error(), msgIfErr)
		}

		if os.Getenv(ENV_OP_TRACE_DURATION) == "true" {

			l.SetDuration(time.Since(start))
		}

		this.AccessLogger.PushTraceLogs(ctx, l)
	}
}

func (this *OperationLogger) PushTraceError(op string, err error, defaultMsg string, ctx context.Context) {

	if !this.CouldLog(ctx) {

		return
	}

	this.AccessLogger.PushTraceLogs(
		ctx,
		logPattern{
			Operation: op,
			Message:   libCommon.Ternary(defaultMsg == "", err.Error(), defaultMsg),
		},
	)
}

func empty_trace_func(err error) {

}
