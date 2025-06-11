package opLog

import (
	"app/internal/debug"
	libCommon "app/internal/lib/common"
	"context"
)

type (
	DebugLogger struct {
		general_logger_t
	}
)

func (this *DebugLogger) newDebug(logUnit string) ILogUseCase {

	clone := libCommon.PointerPrimitive(*this)

	clone.logUnit = logUnit

	return clone
}

func (this *DebugLogger) Debug(logUnit string) ILogUseCase {

	return this.newDebug(logUnit)
}

func (this *DebugLogger) PushCustom(ctx context.Context, lines ...interface{}) {

	this.general_logger_t.pushArbitrary(LOG_LEVEL_DEBUG, ctx, lines)
}

func (this *DebugLogger) isDebugContext(ctx context.Context) bool {

	switch v := (ctx.Value(debug.DEBUG_CONTEXT_KEY)).(type) {
	case bool:
		return v
	default:
		return false
	}
}

func (this *DebugLogger) couldLog(ctx context.Context) bool {

	// return os.Getenv(bootstrap.ENV_DEBUG_LOG) == "true" || this.isDebugContext(ctx)

	return debug.IsDebugging() || this.isDebugContext(ctx)
}

func (this *DebugLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	if !this.couldLog(ctx) {

		return empty_trace_func
	}

	return this.general_logger_t.Messure(LOG_LEVEL_DEBUG, op, msg, ctx)
}

func (this *DebugLogger) PushIfError(err error, op string, msg string, ctx context.Context) {

	if !this.couldLog(ctx) {

		return
	}

	this.general_logger_t.PushIfError(LOG_LEVEL_DEBUG, err, op, msg, ctx)
}

func (this *DebugLogger) Push(op string, msg string, ctx context.Context) {

	if !this.couldLog(ctx) {

		return
	}

	this.general_logger_t.Push(LOG_LEVEL_DEBUG, op, msg, ctx)
}

func (this *DebugLogger) PushCond(
	op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	if !this.couldLog(ctx) {

		return empty_push_cond_func
	}

	return this.general_logger_t.PushCond(LOG_LEVEL_DEBUG, op, msgIfNoErr, ctx)
}

func (this *DebugLogger) PushCondWithMessurement(
	op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	if !this.couldLog(ctx) {

		return empty_push_cond_with_messurement_func
	}

	return this.general_logger_t.PushCondWithMessurement(LOG_LEVEL_DEBUG, op, ctx)
}

func (this *DebugLogger) PushError(op string, err error, defaultMsg string, ctx context.Context) {

	if !this.couldLog(ctx) {

		return
	}

	this.general_logger_t.PushError(LOG_LEVEL_DEBUG, op, err, defaultMsg, ctx)
}
