package promise

import (
	"context"
	"sync"
)

type (
	RejectFunc         func(error) bool
	ResolveFunc[T any] func(T) bool

	Promise[T any] interface {
		//Future() future.Future[interface{}]
		StillHolds() bool
		Await() (v T, err error, ok bool)
	}
)

type (
	promise[T any] struct {
		ctx      context.Context
		cancelFn context.CancelCauseFunc
		ch       chan T
		mu       sync.Mutex
		futures  []chan interface{}
	}
)

func newPromiseObject[T any](ctx context.Context) *promise[T] {

	ret := new(promise[T])

	switch {
	case ctx == nil:
		ret.ctx, ret.cancelFn = context.WithCancelCause(context.TODO())
	default:
		ret.ctx, ret.cancelFn = context.WithCancelCause(ctx)
	}

	ret.ch = make(chan T, 1)

	return ret
}

func resolveFn[T any](p *promise[T]) ResolveFunc[T] {

	return func(v T) bool { return p.resolve(v) }
}

func rejectFn[T any](p *promise[T]) RejectFunc {

	return func(e error) bool { return p.reject(e) }
}

func withResolver[T any](ctx context.Context) (*promise[T], ResolveFunc[T], RejectFunc) {

	ret := newPromiseObject[T](ctx)

	return ret, resolveFn(ret), rejectFn(ret)
}

func New[T any](ctx context.Context) (Promise[T], ResolveFunc[T], RejectFunc) {

	return withResolver[T](ctx)
}

func (this *promise[T]) resolve(v T) bool {

	if !this.StillHolds() {

		return false
	}

	defer close(this.ch)
	this.ch <- v

	go this.reply(v)

	this.cancelFn = nil

	return true
}

func (this *promise[T]) reject(err error) bool {

	if !this.StillHolds() {

		return false
	}

	this.cancelFn(err)
	this.cancelFn = nil

	go this.reply(err)

	return true
}

func (this *promise[T]) Await() (val T, err error, ok bool) {

	if !this.StillHolds() {

		ok = false
		return
	}

	select {
	case val, ok = <-this.ch:
	case <-this.ctx.Done():
	}

	switch {
	case this.ctx.Err() != nil:
		err = this.ctx.Err()
		return
	default:
		return
	}
}

func (this *promise[T]) StillHolds() bool {

	return this.cancelFn != nil
}

func (this *promise[T]) _future() PromiseFuture {

	ret := make(chan interface{}, 1)

	if !this.StillHolds() {

		close(ret)
		return ret
	}

	this.mu.Lock()
	defer this.mu.Unlock()

	this.futures = append(this.futures, ret)

	return ret
}

func (this *promise[T]) reply(v interface{}) {

	for _, fu := range this.futures {

		fu <- v
		close(fu)
	}

	this.futures = nil
}
