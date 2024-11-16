package cacheListService

import (
	"app/internal/bootstrap"
	"app/unitOfWork"
	"context"
	"os"
)

const (
	ENV_CACHE_LOG = bootstrap.ENV_CACHE_LOG
)

type (
	CacheListLogger struct {
		unitOfWork.OperationLogger
	}
)

func (this *CacheListLogger) IsCacheLogEnabled() bool {

	return os.Getenv(ENV_CACHE_LOG) == "true"
}

func (this *CacheListLogger) Messure(op string, msg string, ctx context.Context) func(err error) {

	if !this.IsCacheLogEnabled() {

		ctx = nil
	}

	return this.OperationLogger.Messure(op, msg, ctx)
}

func (this *CacheListLogger) PushTraceIfError(err error, op string, msg string, ctx context.Context) {

	if !this.IsCacheLogEnabled() {

		return
	}

	this.OperationLogger.PushTraceIfError(err, op, msg, ctx)
}

func (this *CacheListLogger) PushTrace(op string, msg string, ctx context.Context) {

	if !this.IsCacheLogEnabled() {

		return
	}

	this.OperationLogger.PushTrace(op, msg, ctx)
}

func (this *CacheListLogger) PushTraceCond(op string, msgIfNoErr string, ctx context.Context) (logErrFunc func(err error, msgIfErr string)) {

	if !this.IsCacheLogEnabled() {

		ctx = nil
	}

	return this.OperationLogger.PushTraceCond(op, msgIfNoErr, ctx)
}

func (this *CacheListLogger) PushTraceCondWithMessurement(
	op string, ctx context.Context,
) func(msgIfNoErr string, err error, msgIfErr string) {

	if !this.IsCacheLogEnabled() {

		ctx = nil
	}

	return this.OperationLogger.PushTraceCondWithMessurement(op, ctx)
}

func (this *CacheListLogger) PushTraceError(op string, err error, defaultMsg string, ctx context.Context) {

	if !this.IsCacheLogEnabled() {

		ctx = nil
	}

	this.OperationLogger.PushTraceError(op, err, defaultMsg, ctx)
}
