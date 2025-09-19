package annotate

import (
	"app/shared/common/stack"
	"reflect"

	"golang.org/x/exp/maps"
)

var (
	singletons            map[reflect.Type]reflect.Value
	singleton_layer_stack stack.Stack[map[reflect.Type]reflect.Value]
)

func tryCleanSingleton() {

	if effector_stack.IsEmpty() && accumulator_stack.IsEmpty() {

		maps.Clear(singletons)
		singletons = nil
	}
}

func StackSingletonLayer() {

	if singletons == nil {

		singletons = make(map[reflect.Type]reflect.Value)
	}

	singleton_layer_stack.Push(singletons)

	singletons = maps.Clone(singletons)
}

func PopSingletonLayer() {

	singletons, _ = singleton_layer_stack.Pop()
}

func GetAcknowledgedSingletons() map[reflect.Type]reflect.Value {

	return maps.Clone(singletons)
}

func AcknowledgeSingleton(t reflect.Type, v reflect.Value) {

	singletons[t] = v
}

func GetSingleton(t reflect.Type) (v reflect.Value, ok bool) {

	v, ok = singletons[t]
	return
}

func HasSingleton(t reflect.Type) bool {

	_, ok := singletons[t]

	return ok
}
