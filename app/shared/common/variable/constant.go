package variable

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

type (
	const_t[T any] struct {
		v T
		_ NoCopy
	}
)

type (
	Constant[T any] *const_t[T]
)

func NewConstain[T any](val T) Constant[T] {

	return Constant[T](&const_t[T]{v: val})
}

func ValueOfConst[T any](variable Constant[T]) T {

	return (*variable).Value().(T)
}

func (this *const_t[T]) String() string {

	return fmt.Sprintf(`< [Constain] %T %v >`, this.v, this.v)
}

/*
return the copy of the actual value in order not to change the stored value
*/
func (this *const_t[T]) Value() interface{} {

	if this == nil {

		this = new(const_t[T])
	}

	switch any(this.v).(type) {
	// alsway copy
	case int8, uint8, int16, uint16, int32,
		uint32, int64, uint64, int, uint,
		uintptr, float32, float64,
		complex64, complex128,
		string:
		return this.v
	default:
		return reflect.ValueOf(this.v).Interface()
	}
}

// func Equal[T comparable](var1 InAssignable[T], var2 InAssignable[T]) bool {

// 	return var1.v == var2.v
// }

func CompareConstain[T constraints.Ordered](left Constant[T], right Constant[T]) int {

	switch {
	case left.v > right.v:
		return 1
	case left.v < right.v:
		return -1
	default:
		return 0
	}
}
