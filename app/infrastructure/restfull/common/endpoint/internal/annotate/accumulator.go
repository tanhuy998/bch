package annotate

import (
	"app/shared/common/stack"
)

type (
	IAccumulator interface {
		Accumulate(interface{}) interface{}
		GetAccumulatorKey() interface{}
	}
)

var (
	accumulator_stack stack.Stack[IAccumulator]
)

func StackAccumulator(v IAccumulator) {

	accumulator_stack.Push(v)
}

func PopAccumulator() {

	accumulator_stack.Pop()

	tryCleanSingleton()
}

func GetAccumulatorQueue() []IAccumulator {

	return accumulator_stack.Snapshot()
}
