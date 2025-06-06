package opLog

import (
	libCommon "app/internal/lib/common"
	accessLogServicePort "app/port/accessLog"
	"context"
	"os"
	"time"
)

type (
	/*
		Default logger checks for the AccessLogger loggable capability of a specific context
		and then push log to the context.
	*/
	default_logger_t struct {
		AccessLogger accessLogServicePort.IAccessLogger
		logUnit      string
	}
)

func (this *default_logger_t) couldLog(ctx context.Context) bool {

	return ctx != nil && this.AccessLogger.IsLogging(ctx) && this.AccessLogger.IsTraceLogging(ctx)
}

func (this *default_logger_t) isAccessLoggerCompatibale(ctx context.Context) bool {

	return this.couldLog(ctx)
}

func (this *default_logger_t) pushArbitrary(level string, ctx context.Context, lines []interface{}) {

	if !this.couldLog(ctx) {

		return
	}

	for _, log := range lines {

		this.AccessLogger.PushTraceLogs(
			ctx,
			newCustomLogPattern(
				logPattern{
					LogLevel:   level,
					LogContext: LOG_CONTEXT_DEFAULT,
					LogUnit:    this.logUnit,
				},
				log,
			),
		)
	}
}

func (this *default_logger_t) messure(
	level string, op string, msg string, ctx context.Context,
) func(err error) {

	start := time.Now()

	return func(err error) {

		if !this.couldLog(ctx) {

			return
		}

		duration := time.Since(start)

		l := &logPattern{
			Operation:  op,
			Message:    msg,
			LogLevel:   level,
			LogContext: LOG_CONTEXT_DEFAULT,
			LogUnit:    this.logUnit,
		}

		l.SetDuration(duration)

		this.AccessLogger.PushTraceLogs(ctx, l)
	}
}

func (this *default_logger_t) pushIfError(
	level string, err error, op string, msg string, ctx context.Context,
) {

	if err == nil || !this.couldLog(ctx) {

		return
	}

	this.AccessLogger.PushTraceLogs(
		ctx,
		logPattern{
			Operation:  op,
			Message:    msg,
			LogLevel:   level,
			ErrorMsg:   err.Error(),
			LogContext: LOG_CONTEXT_DEFAULT,
			LogUnit:    this.logUnit,
		},
	)
}

func (this *default_logger_t) push(
	level string, op string, msg string, ctx context.Context,
) {

	if !this.couldLog(ctx) {

		return
	}

	this.AccessLogger.PushTraceLogs(
		ctx,
		logPattern{
			Operation:  op,
			Message:    msg,
			LogLevel:   level,
			LogContext: LOG_CONTEXT_DEFAULT,
			LogUnit:    this.logUnit,
		},
	)
}

func (this *default_logger_t) pushCond(
	level string, op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	return func(err error, errMsg string) {

		if !this.couldLog(ctx) {

			return
		}

		if err == nil {

			this.push(level, op, msgIfNoErr, ctx)
			return
		}

		this.pushError(level, op, err, msgIfNoErr, ctx)
	}
}

func (this *default_logger_t) pushCondWithMessurement(
	level string, op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	start := time.Now()

	return func(msgIfNoErr string, err error, msgIfErr string) {

		if !this.couldLog(ctx) {

			return
		}

		l := logPattern{
			Operation:  op,
			LogLevel:   level,
			LogContext: LOG_CONTEXT_DEFAULT,
			LogUnit:    this.logUnit,
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

func (this *default_logger_t) pushError(
	level string, op string, err error, defaultMsg string, ctx context.Context,
) {

	if !this.couldLog(ctx) {

		return
	}

	this.AccessLogger.PushTraceLogs(
		ctx,
		logPattern{
			Operation:  op,
			Message:    libCommon.Ternary(defaultMsg == "", err.Error(), defaultMsg),
			LogLevel:   level,
			LogContext: LOG_CONTEXT_DEFAULT,
			LogUnit:    this.logUnit,
		},
	)
}
