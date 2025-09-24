package libCommon

import (
	"reflect"
)

func DereferenceValueOf[Expected_Type_T any](p interface{}) (ret Expected_Type_T, ok bool) {

	var temp interface{} = p

	v := reflect.ValueOf(temp)

	for v.Kind() == reflect.Pointer {

		temp = v.Elem()
		v = reflect.ValueOf(temp)
	}

	ret, ok = temp.(Expected_Type_T)
	return
}

func EvaluateIndirectValueOfPointer[Pointer_T comparable](indirectPtr *Pointer_T) {

	if indirectPtr == nil {

		panic("bad arg passed to ptr, nil given")
	}

	t := reflect.TypeFor[Pointer_T]()

	switch {
	case t.Kind() != reflect.Pointer:
		panic("the type parameter Pointer_T is not kind of pointer")
	}

	EvaluateIndirectValueBy(
		reflect.Indirect(
			reflect.ValueOf(indirectPtr),
		),
	)
}

func EvaluateIndirectValueOf[T any](ptr *T) {

	if ptr == nil {

		panic("bad arg passed to ptr, nil given")
	}

	EvaluateIndirectValueBy(
		reflect.Indirect(
			reflect.ValueOf(ptr),
		),
	)
}

/*
Indirectly evaluate zero value by a reflecting the pointer which pointed
to a specific variable (a memory location).
*/
func EvaluateIndirectValueBy(ptr reflect.Value) {

	switch {
	case ptr.Type().Kind() != reflect.Pointer:
		panic("type of reflection value is not pointer")
	case !ptr.IsNil():
		return
	default:
		ptr.Set(
			reflect.New(ptr.Type().Elem()),
		)
	}
}
