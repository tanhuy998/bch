package sync

import (
	"context"
	"sync"
	"time"
)

type (
	NullableContext[T context.Context] struct {
		mutex       sync.Mutex
		baseContext *T //context.Context
	}
)

func NewNullableContext[T context.Context](baseCtx *T) *NullableContext[T] {

	return &NullableContext[T]{
		baseContext: baseCtx,
	}
}

func (this *NullableContext[T]) waitForBaseContextEvaluated() {

	if this.baseContext != nil {

		return
	}

	this.mutex.Lock()

	for this.baseContext == nil {

	}

	this.mutex.Unlock()
}

func (this *NullableContext[T]) GetBaseContext() *T {

	this.waitForBaseContextEvaluated()

	return this.baseContext
}

func (this *NullableContext[T]) Deadline() (deadline time.Time, ok bool) {

	return (*this.GetBaseContext()).Deadline()
}

func (this *NullableContext[T]) Done() <-chan struct{} {

	return (*this.GetBaseContext()).Done()
}

func (this *NullableContext[T]) Err() error {

	return (*this.GetBaseContext()).Err()
}

func (this *NullableContext[T]) Value(key any) any {

	return (*this.GetBaseContext()).Value(key)
}

func (this *NullableContext[T]) SetBaseContext(ctx *T) {

	this.baseContext = ctx

	this.waitForBaseContextEvaluated()
}
