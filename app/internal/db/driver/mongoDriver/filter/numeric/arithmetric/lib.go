package arithmetric

import "app/internal/db/driver/mongoDriver/filter/chain"

func newChainNode(operator string, left interface{}, right interface{}) *chain.ExpressionChainNode {

	ret := new(chain.ExpressionChainNode)

	ret.SetOperator(operator)
	ret.SetLeftOperand(left)
	ret.SetRightOperand(right)

	return ret
}

func Add(left interface{}, right interface{}) *chain.ExpressionChainNode {

	return newChainNode("$add", left, right)
}

func Substract(left interface{}, right interface{}) *chain.ExpressionChainNode {

	return newChainNode("$substract", left, right)
}

func Multiply(left interface{}, right interface{}) *chain.ExpressionChainNode {

	return newChainNode("$multiply", left, right)
}

func Divide(left interface{}, right interface{}) *chain.ExpressionChainNode {

	return newChainNode("$divide", left, right)
}

func Power(left interface{}, right interface{}) *chain.ExpressionChainNode {

	return newChainNode("$pow", left, right)
}

func Logarit(left interface{}, base interface{}) *chain.ExpressionChainNode {

	return newChainNode("$log", left, base)
}

// func Sqrt(operand interface{}) *chain.ExpressionChainNode {

// 	return newChainNode("$sqrt", left,)
// }
