package opLog

import (
	"context"
)

type (
	/*
		General logger is a strategic log dispatcher that accepts any passed context
		and send the received log message to compatible context logger
	*/
	general_logger_t struct {
		adaptive_logger_t
	}
)

func (this *general_logger_t) General() IGeneralLogger {

	return this
}

func (this *general_logger_t) PushCustom(level string, ctx context.Context, lines ...interface{}) {

	this.pushArbitrary(level, ctx, lines)
}

func (this *general_logger_t) pushArbitrary(level string, ctx context.Context, lines []interface{}) {

	this.defaultStrategyFor(ctx).pushArbitrary(level, ctx, lines)
}

func (this *general_logger_t) defaultStrategyFor(ctx context.Context) IInternalGeneralLogger {

	logRecorder, isLogRecorder := this.adaptive_logger_t.isLogRecorder(ctx)

	accessLogCompatible := this.default_logger_t.isAccessLoggerCompatibale(ctx)

	if accessLogCompatible {

		return &this.default_logger_t
	}

	if isLogRecorder && this.tryMergeLogs(logRecorder) {

		return &this.default_logger_t

	} else if isLogRecorder {

		return &this.adaptive_logger_t
	}

	return &this.default_logger_t
}

func (this *general_logger_t) TryMergeLogs(ctx context.Context) bool {

	switch logCtx, isLogCtx := this.adaptive_logger_t.isLogRecorder(ctx); {
	case isLogCtx:
		return this.tryMergeLogs(logCtx)
	default:
		return false
	}
}

func (this *general_logger_t) tryMergeLogs(logCtx ILogRecorderContext) bool {

	if logCtx == nil {

		return false
	}

	baseCtx := logCtx.GetBaseContext()

	if !this.default_logger_t.isAccessLoggerCompatibale(baseCtx) {

		return false
	}

	for line := range logCtx.Release() {

		this.default_logger_t.AccessLogger.PushTraceLogs(baseCtx, line)
	}

	return true
}

func (this *general_logger_t) Messure(
	level string, op string, msg string, ctx context.Context,
) func(err error) {

	return this.defaultStrategyFor(ctx).
		messure(level, op, msg, ctx)
}

func (this *general_logger_t) PushIfError(
	level string, err error, op string, msg string, ctx context.Context,
) {

	this.defaultStrategyFor(ctx).
		pushIfError(level, err, op, msg, ctx)
}

func (this *general_logger_t) Push(
	level string, op string, msg string, ctx context.Context,
) {

	this.defaultStrategyFor(ctx).
		push(level, op, msg, ctx)
}

func (this *general_logger_t) PushCond(
	level string, op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	return this.defaultStrategyFor(ctx).
		pushCond(level, op, msgIfNoErr, ctx)
}

func (this *general_logger_t) PushCondWithMessurement(
	level string, op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	return this.defaultStrategyFor(ctx).
		pushCondWithMessurement(level, op, ctx)
}

func (this *general_logger_t) PushError(
	level string, op string, err error, defaultMsg string, ctx context.Context,
) {

	this.defaultStrategyFor(ctx).
		pushError(level, op, err, defaultMsg, ctx)
}
