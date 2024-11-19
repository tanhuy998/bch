package opLog

import (
	libCommon "app/internal/lib/common"
	accessLogServicePort "app/port/accessLog"
	"context"
	"os"
	"time"
)

const (
	LOG_LEVEL_TRACE = "trace"
	LOG_LEVEL_DEBUG = "debug"
)

func empty_trace_func(err error) {}

func empty_push_cond_func(err error, msgIfErr string) {}

func empty_push_cond_with_messurement_func(msgIfNoErr string, err error, msgIfErr string) {}

func couldLog(logger accessLogServicePort.IAccessLogger, ctx context.Context) bool {

	return ctx != nil && logger.IsLogging(ctx) && logger.IsTraceLogging(ctx)
}

func messure(
	logger accessLogServicePort.IAccessLogger, level string, op string, msg string, ctx context.Context,
) func(err error) {

	start := time.Now()

	return func(err error) {

		if !couldLog(logger, ctx) {

			return
		}

		duration := time.Since(start)

		l := &logPattern{
			Operation: op,
			Message:   msg,
			LogLevel:  level,
		}

		l.SetDuration(duration)

		logger.PushTraceLogs(ctx, l)
	}
}

func pushTraceIfError(logger accessLogServicePort.IAccessLogger, level string, err error, op string, msg string, ctx context.Context) {

	if err == nil || !couldLog(logger, ctx) {

		return
	}

	logger.PushTraceLogs(
		ctx,
		logPattern{
			Operation: op,
			Message:   msg,
			LogLevel:  level,
			ErrorMsg:  err.Error(),
		},
	)
}

func pushTrace(
	logger accessLogServicePort.IAccessLogger, level string, op string, msg string, ctx context.Context,
) {

	if !couldLog(logger, ctx) {

		return
	}

	logger.PushTraceLogs(
		ctx,
		logPattern{
			Operation: op,
			Message:   msg,
			LogLevel:  level,
		},
	)
}

func pushTraceCond(
	logger accessLogServicePort.IAccessLogger, level string, op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	return func(err error, errMsg string) {

		if !couldLog(logger, ctx) {

			return
		}

		if err == nil {

			pushTrace(logger, level, op, msgIfNoErr, ctx)
			return
		}

		pushTraceError(logger, level, op, err, msgIfNoErr, ctx)
	}
}

func PushTraceCondWithMessurement(
	logger accessLogServicePort.IAccessLogger, level string, op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	start := time.Now()

	return func(msgIfNoErr string, err error, msgIfErr string) {

		if !couldLog(logger, ctx) {

			return
		}

		l := logPattern{
			Operation: op,
			LogLevel:  level,
		}

		if err == nil {

			l.Message = msgIfNoErr

		} else {

			l.Message = libCommon.Ternary(msgIfErr == "", err.Error(), msgIfErr)
		}

		if os.Getenv(ENV_OP_TRACE_DURATION) == "true" {

			l.SetDuration(time.Since(start))
		}

		logger.PushTraceLogs(ctx, l)
	}
}

func pushTraceError(
	logger accessLogServicePort.IAccessLogger, level string, op string, err error, defaultMsg string, ctx context.Context,
) {

	if !couldLog(logger, ctx) {

		return
	}

	logger.PushTraceLogs(
		ctx,
		logPattern{
			Operation: op,
			Message:   libCommon.Ternary(defaultMsg == "", err.Error(), defaultMsg),
			LogLevel:  level,
		},
	)
}
