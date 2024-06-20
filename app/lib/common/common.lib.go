package libCommon

import (
	"fmt"
	"reflect"
)

func Ternary[T any](criteria bool, valIfTrue T, valIfFalse T) T {

	if criteria {
		fmt.Println("true", criteria)
		return valIfTrue
	}

	return valIfFalse
}

func PointerPrimitive[T any](val T) *T {

	ret := val
	return &ret
}

func Or(expressions ...bool) bool {

	for _, isTrue := range expressions {

		if isTrue {

			return true
		}
	}

	return false
}

func GetOriginalTypeOf(value interface{}) reflect.Type {

	t := reflect.TypeOf(value)
	for t.Kind() == reflect.Ptr {

		t = t.Elem()
	}

	return t
}

func GetOriginalInterfaceOf(value interface{}) reflect.Type {

	t := reflect.TypeOf(value)
	for t.Kind() == reflect.Ptr {

		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("called libCommon.GetOriginalTypeOf with a value that is not a pointer to an interface. (*MyInterface)(nil)")
	}
	return t
}

func Wrap[T any]() reflect.Type {

	return GetOriginalTypeOf((*T)(nil))
}

func IsInterface[T any]() bool {

	return reflect.TypeOf((*T)(nil)).Kind() == reflect.Interface
}

func ReverseSlice[T any](list ...T) []T {

	var left int = 0
	var right int = len(list) - 1

	for left < right {

		temp := list[left]
		list[left] = list[right]
		list[right] = temp

		left++
		right--
	}

	return list
}
