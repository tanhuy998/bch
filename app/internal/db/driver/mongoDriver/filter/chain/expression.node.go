package chain

import (
	"app/internal/db/driver/mongoDriver/expression"
)

type (
	ExpressionChainNode struct {
		//iterator.Iterator
		expression.BinaryExpression[interface{}, interface{}]
	}
)

func (this ExpressionChainNode) Next() *ExpressionChainNode {

	switch v := this.GetLeftOperand().(type) {
	case *ExpressionChainNode:
		return v
	default:
		return nil
	}
}

func (this ExpressionChainNode) Left() interface{} {

	return this.GetLeftOperand()
}

func (this ExpressionChainNode) Right() interface{} {

	return this.GetRightOperand()
}
