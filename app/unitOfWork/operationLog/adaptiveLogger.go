package opLog

import (
	libCommon "app/internal/lib/common"
	"context"
	"time"
)

type (
	ILogRecorderContext interface {
		Release() <-chan interface{}
		GetBaseContext() context.Context
		SetBaseContext(ctx context.Context)
		PushLog(lines interface{})
	}
)

type (
	/*
		Adaptive logger checks if a given context is log recorder that await (or has already wrap)
		a compatible AccessLogger context.
	*/
	adaptive_logger_t struct {
		default_logger_t
	}
)

func (this *adaptive_logger_t) isLogRecorder(ctx context.Context) (ILogRecorderContext, bool) {

	ret, ok := ctx.(ILogRecorderContext)

	return ret, ok
}

func (this *adaptive_logger_t) pushArbitrary(level string, ctx context.Context, lines []interface{}) {

	logCtx, isLogCtx := this.isLogRecorder(ctx)

	if !isLogCtx {
		if !this.couldLog(ctx) {

			return
		}

		return
	}

	for _, customLog := range lines {

		logCtx.PushLog(
			newCustomLogPattern(
				logPattern{
					LogLevel:   level,
					LogContext: LOG_CONTEXT_ADAPTIVE,
					LogUnit:    this.logUnit,
				},
				customLog,
			),
		)
	}
}

func (this *adaptive_logger_t) couldLog(ctx context.Context) bool {

	_, ok := this.isLogRecorder(ctx)

	return ok
}

func (this *adaptive_logger_t) messure(
	level string, op string, msg string, ctx context.Context,
) func(err error) {

	debugCtx, ok := this.isLogRecorder(ctx)

	if !ok {

		return empty_trace_func
	}

	start := time.Now()

	return func(err error) {

		line := newAdaptiveContextLogPattern(level, op, msg)
		line.LogUnit = this.logUnit

		if err != nil {

			line.ErrorMsg = err.Error()
		}

		line.SetDuration(time.Since(start))

		debugCtx.PushLog(
			*line,
		)
	}
}

func (this *adaptive_logger_t) pushIfError(
	level string, err error, op string, msg string, ctx context.Context,
) {

	if err == nil {

		return
	}

	debugCtx, ok := this.isLogRecorder(ctx)

	if !ok {

		return
	}

	line := newAdaptiveContextLogPattern(level, op, msg)

	line.LogUnit = this.logUnit
	line.ErrorMsg = err.Error()

	debugCtx.PushLog(*line)
}

func (this *adaptive_logger_t) push(
	level string, op string, msg string, ctx context.Context,
) {

	debugCtx, ok := this.isLogRecorder(ctx)

	if !ok {

		return
	}

	line := newAdaptiveContextLogPattern(level, op, msg)
	line.LogUnit = this.logUnit

	debugCtx.PushLog(*line)
}

func (this *adaptive_logger_t) pushCond(
	level string, op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	debugCtx, ok := this.isLogRecorder(ctx)

	if !ok {

		return empty_push_cond_func
	}

	return func(err error, msgIfErr string) {

		msg := libCommon.Ternary(err != nil, msgIfErr, msgIfNoErr)

		line := newAdaptiveContextLogPattern(level, op, msg)
		line.LogUnit = this.logUnit

		if err != nil {

			line.ErrorMsg = err.Error()
		}

		debugCtx.PushLog(*line)
	}
}

func (this *adaptive_logger_t) pushCondWithMessurement(
	level string, op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	debugCtx, ok := this.isLogRecorder(ctx)

	if !ok {

		return empty_push_cond_with_messurement_func
	}

	start := time.Now()

	return func(msgIfNoErr string, err error, msgIfErr string) {

		msg := libCommon.Ternary(err != nil, msgIfErr, msgIfNoErr)

		line := newAdaptiveContextLogPattern(level, op, msg)
		line.LogUnit = this.logUnit

		if err != nil {

			line.ErrorMsg = err.Error()
		}

		line.SetDuration(time.Since(start))

		debugCtx.PushLog(*line)
	}
}

func (this *adaptive_logger_t) pushError(
	level string, op string, err error, defaultMsg string, ctx context.Context,
) {

	switch debugCtx, ok := this.isLogRecorder(ctx); {
	case !ok:
		return
	case err == nil:
		return
	default:
		line := newAdaptiveContextLogPattern(level, op, defaultMsg)
		line.ErrorMsg = err.Error()
		line.LogUnit = this.logUnit

		debugCtx.PushLog(*line)

		return
	}
}
