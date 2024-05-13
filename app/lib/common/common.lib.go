package libCommon

import "reflect"

func Ternary[T any](criteria bool, valIfTrue T, valIfFalse T) T {

	if criteria {

		return valIfTrue
	}

	return valIfFalse
}

func PointerPrimitive[T any](val T) *T {

	ret := val
	return &ret
}

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)
	for t.Kind() == reflect.Ptr {

		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("called inject.InterfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}
	return t
}
