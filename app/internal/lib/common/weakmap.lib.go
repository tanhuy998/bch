package libCommon

import (
	"unsafe"
)

type (
	weak_map_t[Key_T comparable] map[Key_T]unsafe.Pointer
)

func (this *weak_map_t[Key_T]) init() {

	if *this == nil {

		*this = make(map[Key_T]unsafe.Pointer)
	}
}

type (
	WeakMap[Key_T comparable, Value_T any] struct {
		m weak_map_t[Key_T]
	}
)

func (this *WeakMap[Key_T, Value_T]) Set(key Key_T, val Value_T) {

	this.m.init()

	this.m[key] = unsafe.Pointer(&val)
}

func (this *WeakMap[Key_T, Value_T]) Get(key Key_T) (val Value_T, ok bool) {

	if this.m == nil {

		return
	}

	this.m.init()

	switch weak, _ok := this.m[key]; {
	case !_ok:
		return
	default:
		p := (*Value_T)(weak)

		if p == nil {

			return
		}

		val = *p
		return
	}
}

func (this *WeakMap[Key_T, Value_T]) Delete(key Key_T) {

	switch {
	case this.m == nil:
		return
	default:
		this.m.init()
	}

	delete(this.m, key)
}
