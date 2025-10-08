package objectPool

import (
	libCommon "app/internal/lib/common"
	"app/internal/lib/sync"
	"reflect"
)

type (
	key_constraint[T any] struct{}
)

var (
	/*
		Global object manipulator
	*/
	pool_map libCommon.WriteConcernMap[reflect.Type, interface{}]
)

func LoadPoolFor[T any]() *sync.Recycler[T] {

	//return __resolve_pool__[T](pool_map[reflect.TypeFor[T]()])

	switch v, ok := pool_map.Load(reflect.TypeFor[T]()); {
	case ok:
		return __resolve_pool[T](v)
	default:
		return nil
	}
}

func LoadOrInitPoolFor[T any]() *sync.Recycler[T] {

	key := reflect.TypeFor[key_constraint[T]]()

	switch reflectValue, ok := pool_map.Load(key); {
	case !ok:
		return __init_for[T]()
	default:
		ret := __resolve_pool[T](reflectValue)

		if ret == nil {

			return __init_for[T]()
		}

		return ret
	}
}

func __init_for[T any]() *sync.Recycler[T] {

	key := reflect.TypeFor[key_constraint[T]]()

	pool := new(sync.Recycler[T])
	//pool_map[key] = reflect.ValueOf(pool)

	pool_map.Store(key, pool)

	return pool
}

func __resolve_pool[T any](in interface{}) *sync.Recycler[T] {

	switch ret := in.(type) {
	case *sync.Recycler[T]:
		return ret
	default:
		return nil
	}
}
