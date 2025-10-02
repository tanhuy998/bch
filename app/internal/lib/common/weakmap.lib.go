package libCommon

import (
	"app/shared/common/weak"
)

type (
	weak_map_t[Key_T comparable, Value_T any] map[Key_T]weak.Pointer[Value_T] // unsafe.Pointer
)

func (this *weak_map_t[Key_T, Value_T]) init() {

	if *this == nil {

		*this = make(map[Key_T]weak.Pointer[Value_T])

		__watch_weak_map__(*this)
	}
}

func (this weak_map_t[Key_T, Value_T]) clean() {

	if this == nil {

		return
	}

	for key, weakPointer := range this {

		switch weakPointer.Value() {
		case nil:
			delete(this, key)
		}
	}
}

func (this *weak_map_t[Key_T, Value_T]) set(key Key_T, val Value_T) {

	this.init()

	(*this)[key] = weak.Make(&val)
}

func (this *weak_map_t[Key_T, Value_T]) get(key Key_T) (v Value_T, ok bool) {

	if *this == nil {

		return
	}

	switch weakPointer, exists := (*this)[key]; {
	case !exists:
		return
	default:
		pointer := weakPointer.Value()

		if pointer == nil {
			delete(*this, key)
			return
		}

		return *pointer, true
	}
}

type (
	// WeakMap store values by keys, value that strored in weak map
	// could be collected by garbage collector.
	// WeakMap is not designed to be thread safe and race safety.
	WeakMap[Key_T comparable, Value_T any] struct {
		m weak_map_t[Key_T, Value_T]
	}
)

func (this *WeakMap[Key_T, Value_T]) Set(key Key_T, val Value_T) {

	this.m.set(key, val)
}

func (this *WeakMap[Key_T, Value_T]) Get(key Key_T) (val Value_T, ok bool) {

	return this.m.get(key)
}

func (this *WeakMap[Key_T, Value_T]) Delete(key Key_T) {

	switch {
	case this.m == nil:
		return
	default:
		delete(this.m, key)
	}
}
