package opLog

import (
	"app/internal/bootstrap"
	"context"
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
		DebugLogger
	}
)

func (this *OperationLogger) CouldLog(ctx context.Context) bool {

	return ctx != nil && this.AccessLogger.IsLogging(ctx) && this.AccessLogger.IsTraceLogging(ctx)
}

func (this *OperationLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	return messure(this.AccessLogger, LOG_LEVEL_TRACE, op, msg, ctx)
}

func (this *OperationLogger) PushTraceIfError(err error, op string, msg string, ctx context.Context) {

	pushTraceIfError(this.AccessLogger, LOG_LEVEL_TRACE, err, op, msg, ctx)
}

func (this *OperationLogger) PushTrace(op string, msg string, ctx context.Context) {

	pushTrace(this.AccessLogger, LOG_LEVEL_TRACE, op, msg, ctx)
}

func (this *OperationLogger) PushTraceCond(
	op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	return pushTraceCond(this.AccessLogger, LOG_LEVEL_TRACE, op, msgIfNoErr, ctx)
}

func (this *OperationLogger) PushTraceCondWithMessurement(
	op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	return PushTraceCondWithMessurement(this.AccessLogger, LOG_LEVEL_TRACE, op, ctx)
}

func (this *OperationLogger) PushTraceError(op string, err error, defaultMsg string, ctx context.Context) {

	pushTraceError(this.AccessLogger, LOG_LEVEL_TRACE, op, err, defaultMsg, ctx)
}
