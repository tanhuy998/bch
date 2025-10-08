package objectPool

import (
	"app/internal/lib/sync"
)

type (
	IDisposable interface {
		Dispose()
	}
)

type (
	/*
		Ready to use embedding implementation of internal/objectPool.
	*/
	ObjectRecycler[T any] struct {
		pool *sync.Recycler[T]
	}
)

func _getPresenterPool[T any]() *sync.Recycler[T] {

	return LoadOrInitPoolFor[T]()
}

func (this *ObjectRecycler[T]) init() {

	if this.pool == nil {

		this.pool = _getPresenterPool[T]()
	}
}

func (this *ObjectRecycler[T]) Load() *T {

	this.init()

	switch ret, ok := this.pool.Get(); {
	case !ok, ret == nil:
		return new(T)
	default:
		return ret
	}
}

func (this *ObjectRecycler[T]) Collect(obj *T) {

	if obj == nil {

		return
	}

	this.init()

	this.pool.Recycle(obj)

	switch v := any(obj).(type) {
	case IDisposable:
		v.Dispose()
	}
}

func LoadObjetFor[T any]() *T {

	pool := LoadPoolFor[T]()

	if pool == nil {

		return nil
	}

	ret, _ := pool.Get()

	return ret
}
