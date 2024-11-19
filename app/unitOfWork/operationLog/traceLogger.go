package opLog

import (
	libCommon "app/internal/lib/common"
	accessLogServicePort "app/port/accessLog"
	"context"
	"os"
	"time"
)

type (
	TraceLogger struct {
		AccessLogger accessLogServicePort.IAccessLogger
	}
)

func (this *TraceLogger) CouldLog(ctx context.Context) bool {

	return ctx != nil && this.AccessLogger.IsLogging(ctx) && this.AccessLogger.IsTraceLogging(ctx)
}

func (this *TraceLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

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

func (this *TraceLogger) PushTraceIfError(err error, op string, msg string, ctx context.Context) {

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

func (this *TraceLogger) PushTrace(op string, msg string, ctx context.Context) {

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

func (this *TraceLogger) PushTraceCond(
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

func (this *TraceLogger) PushTraceCondWithMessurement(
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

func (this *TraceLogger) PushTraceError(op string, err error, defaultMsg string, ctx context.Context) {

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
