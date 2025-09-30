package opLog

import (
	"app/internal/lib/sync"

	"context"
)

type (
	ILogRecorder interface {
		GetLogRecorder() *LogRecorder
	}
)

type (
	LogRecorder struct {
		sync.NullableContext[context.Context]
		logQueue sync.Queue[interface{}]
	}
)

func NewLogContext(ctx context.Context) *LogRecorder {

	if ctx == nil {

		ctx = context.TODO()
	}

	ret := &LogRecorder{}

	ret.NullableContext.SetBaseContext(&ctx)

	ret.logQueue.Init()

	return ret
}

func (this *LogRecorder) PushLog(lines interface{}) {

	this.logQueue.Push(lines)
}

func (this *LogRecorder) GetBaseContext() context.Context {

	return *this.NullableContext.GetBaseContext()
}

func (this *LogRecorder) SetBaseContext(ctx context.Context) {

	//this.Context = ctx

	this.NullableContext.SetBaseContext(&ctx)
}

func (this *LogRecorder) Release() <-chan interface{} {

	return this.logQueue.Release()
}
