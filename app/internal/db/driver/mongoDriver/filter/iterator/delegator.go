package iterator

import "app/internal/db/driver/mongoDriver/expression"

type (
	ExpressionDelegator[Left_T any, Right_T any] struct {
		iterator expression.IBinaryExpression[interface{}, interface{}] //*Iterator
		lhs      Left_T
		rhs      Right_T
	}
)

func (this *ExpressionDelegator[Left_T, Right_T]) Init() {

	this.iterator.SetLeftOperand(&this.lhs)
	this.iterator.SetRightOperand(&this.rhs)
}

func (this *ExpressionDelegator[Left_T, Right_T]) Accept(iterator expression.IBinaryExpression[interface{}, interface{}]) {

	switch {
	case iterator == nil:
		panic("iterator for IteratorDelegator must not be nil")
	case iterator == this.iterator:
		return
	default:
		this.iterator = iterator
		this.Init()
	}
}

func (this *ExpressionDelegator[Left_T, Right_T]) SetLeftOperand(lhs Left_T) {

	this.lhs = lhs
}

func (this *ExpressionDelegator[Left_T, Right_T]) SetRightOperand(rhs Right_T) {

	this.rhs = rhs
}

func (this *ExpressionDelegator[Left_T, Right_T]) GetLeftOperand() Left_T {

	return this.lhs
}

func (this *ExpressionDelegator[Left_T, Right_T]) GetRightOperand() Right_T {

	return this.rhs
}
