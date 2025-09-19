package annotate

import (
	"app/shared/common/stack"
	"reflect"
)

var (
	effector_stack stack.Stack[reflect.Value]
)

func StackEffector(v reflect.Value) {

	effector_stack.Push(v)
}

func PopEffector() {

	effector_stack.Pop()

	tryCleanSingleton()
}

func GetEffectorStack() *stack.Stack[reflect.Value] {

	return effector_stack.Clone()
}
