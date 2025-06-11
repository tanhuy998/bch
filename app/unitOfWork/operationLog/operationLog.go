package opLog

import (
	"app/internal/bootstrap"
	libCommon "app/internal/lib/common"
	"context"
)

const (
	ENV_OP_TRACE_DURATION = bootstrap.ENV_OP_TRACE_DURATION
)

type (
	// obsolete interface
	// IOperationLogger interface {
	// 	Messure(op string, msg string, ctx context.Context) func(err error)
	// 	PushTraceIfError(err error, op string, msg string, ctx context.Context)
	// 	PushTrace(op string, msg string, ctx context.Context)
	// 	PushTraceCond(op string, msgIfNoErr string, ctx context.Context) (logErrFunc func(err error, msgIfErr string))
	// 	PushTraceCondWithMessurement(
	// 		op string, msgIfNoErr string, ctx context.Context,
	// 	) func(err error, msgIfErr string)
	// 	PushTraceError(op string, err error, defaultMsg string, ctx context.Context)
	// }

	OperationLogger struct {
		DebugLogger
	}
)

func (this *OperationLogger) newTrace(logUnit string) ILogUseCase {

	clone := libCommon.PointerPrimitive(*this)

	clone.logUnit = logUnit

	return clone
}

func (this *OperationLogger) Trace(logUnit string) ILogUseCase {

	return this.newTrace(logUnit)
}

func (this *OperationLogger) PushCustom(ctx context.Context, lines ...interface{}) {

	this.general_logger_t.pushArbitrary(LOG_LEVEL_TRACE, ctx, lines)
}

func (this *OperationLogger) CouldLog(ctx context.Context) bool {

	return ctx != nil && this.AccessLogger.IsLogging(ctx) && this.AccessLogger.IsTraceLogging(ctx)
}

func (this *OperationLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	// return this.messure(LOG_LEVEL_TRACE, op, msg, ctx)

	return this.general_logger_t.Messure(LOG_LEVEL_TRACE, op, msg, ctx)
}

func (this *OperationLogger) PushIfError(err error, op string, msg string, ctx context.Context) {

	// this.pushTraceIfError(LOG_LEVEL_TRACE, err, op, msg, ctx)

	this.general_logger_t.PushIfError(LOG_LEVEL_TRACE, err, op, msg, ctx)
}

func (this *OperationLogger) PushTrace(op string, msg string, ctx context.Context) {

	// this.pushTrace(LOG_LEVEL_TRACE, op, msg, ctx)

	this.general_logger_t.Push(LOG_LEVEL_TRACE, op, msg, ctx)
}

func (this *OperationLogger) PushCond(
	op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	// return this.pushTraceCond(LOG_LEVEL_TRACE, op, msgIfNoErr, ctx)

	return this.general_logger_t.PushCond(LOG_LEVEL_TRACE, op, msgIfNoErr, ctx)
}

func (this *OperationLogger) PushCondWithMessurement(
	op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	// return this.pushTraceCondWithMessurement(LOG_LEVEL_TRACE, op, ctx)

	return this.general_logger_t.PushCondWithMessurement(LOG_LEVEL_TRACE, op, ctx)
}

func (this *OperationLogger) PushError(op string, err error, defaultMsg string, ctx context.Context) {

	// this.pushTraceError(LOG_LEVEL_TRACE, op, err, defaultMsg, ctx)

	this.general_logger_t.PushError(LOG_LEVEL_TRACE, op, err, defaultMsg, ctx)
}
