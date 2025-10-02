package weak

import (
	"reflect"
	"runtime"
	"unsafe"
)

type (
	// go's version: 1.22.
	// Pollyfill for weak.Pointer that is implemented
	// in go version 1.24
	Pointer[Value_T any] struct {
		p unsafe.Pointer
	}
)

func Make[T any](ptr *T) Pointer[T] {

	ret := Pointer[T]{}

	runtime.SetFinalizer(ptr, func(_ *T) {

		ret.p = unsafe.Pointer(nil)
	})

	return ret
}

func (this Pointer[Value_T]) Value() *Value_T {

	if this.p == nil {

		return nil
	}

	switch ret := (*Value_T)(this.p); {
	case ret == nil,
		reflect.TypeOf(ret) != reflect.TypeFor[*Value_T]():
		return nil
	default:
		return ret
	}
}
