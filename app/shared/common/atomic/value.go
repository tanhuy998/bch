package atomic

import "sync/atomic"

type (
	Value[T any] struct {
		v atomic.Value
	}
)

func (this *Value[T]) CompareAndSwap(old T, new T) {

	this.v.CompareAndSwap(old, new)
}

func (this *Value[T]) Load() T {

	return this.v.Load().(T)
}

func (this *Value[T]) Store(value T) {

	this.v.Store(value)
}

func (this *Value[T]) Swap(value T) (old T) {

	return this.v.Swap(value).(T)
}
