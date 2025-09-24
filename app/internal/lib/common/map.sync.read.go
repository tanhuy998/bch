package libCommon

import "sync"

type (
	// Generic type wrapper for builtin sync.Map package
	ReadConcernMap[Key_T comparable, Value_T comparable] struct {
		m sync.Map
	}
)

func (this *ReadConcernMap[Key_T, Value_T]) CompareAndDelete(key Key_T, old Value_T) (deleted bool) {

	return this.m.CompareAndDelete(key, old)
}

func (this *ReadConcernMap[Key_T, Value_T]) CompareAndSwap(key Key_T, old Value_T, new Value_T) bool {

	return this.m.CompareAndSwap(key, old, new)
}

func (this *ReadConcernMap[Key_T, Value_T]) Delete(Key Key_T) {

	this.m.Delete(Key)
}

func (this *ReadConcernMap[Key_T, Value_T]) Load(key Key_T) (val Value_T, ok bool) {

	v, ok := this.m.Load(key)

	if !ok {

		return
	}

	val, ok = v.(Value_T)
	return
}

func (this *ReadConcernMap[Key_T, Value_T]) LoadAndDelete(key Key_T) (val Value_T, deleted bool) {

	v, ok := this.m.LoadAndDelete(key)

	if !ok {

		return
	}

	val, deleted = v.(Value_T)
	return
}

func (this *ReadConcernMap[Key_T, Value_T]) LoadOrStore(key Key_T, val Value_T) (actual Value_T, loaded bool) {

	v, loaded := this.m.LoadOrStore(key, val)

	actual, _ = v.(Value_T)
	return
}

func (this *ReadConcernMap[Key_T, Value_T]) Range(
	fn func(key Key_T, val Value_T) bool,
) {

	this.m.Range(
		func(key, value any) bool {

			return fn(key.(Key_T), value.(Value_T))
		},
	)
}

func (this *ReadConcernMap[Key_T, Value_T]) Store(key Key_T, val Value_T) {

	this.m.Store(key, val)
}

func (this *ReadConcernMap[Key_T, Value_T]) Swap(key Key_T, val Value_T) (previous Value_T, loaded bool) {

	prev, loaded := this.m.Swap(key, val)

	previous = prev.(Value_T)
	return
}
