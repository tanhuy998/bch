package cacheListService

import (
	libCommon "app/internal/lib/common"
	"app/unitOfWork"
	"context"
	"os"
)

var (
	is_cache_logging bool
)

func init() {

	is_cache_logging = libCommon.Ternary(
		os.Getenv("CACHE_LOG") == "true",
		true, false,
	)
}

type (
	CacheListLogger struct {
		unitOfWork.OperationLogger
	}
)

func (this *CacheListLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	if !is_cache_logging {

		ctx = nil
	}

	return this.OperationLogger.Messure(op, msg, ctx)
}

func (this *CacheListLogger) PushTraceIfError(err error, op string, msg string, ctx context.Context) {

	if !is_cache_logging {

		return
	}

	this.OperationLogger.PushTraceIfError(err, op, msg, ctx)
}

func (this *CacheListLogger) PushTrace(op string, msg string, ctx context.Context) {

	if !is_cache_logging {

		return
	}

	this.OperationLogger.PushTrace(op, msg, ctx)
}

func (this *CacheListLogger) PushTraceCond(op string, msgIfNoErr string, ctx context.Context) (logErrFunc func(err error, msgIfErr string)) {

	if !is_cache_logging {

		ctx = nil
	}

	return this.OperationLogger.PushTraceCond(op, msgIfNoErr, ctx)
}

func (this *CacheListLogger) PushTraceCondWithMessurement(
	op string, msgIfNoErr string, ctx context.Context,
) func(err error, msgIfErr string) {

	if !is_cache_logging {

		ctx = nil
	}

	return this.OperationLogger.PushTraceCondWithMessurement(op, msgIfNoErr, ctx)
}

func (this *CacheListLogger) PushTraceError(op string, err error, defaultMsg string, ctx context.Context) {

	if !is_cache_logging {

		ctx = nil
	}

	this.OperationLogger.PushTraceError(op, err, defaultMsg, ctx)
}
