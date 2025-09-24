package expression

import "go.mongodb.org/mongo-driver/bson"

type (
	VariadicOperandExpression struct {
		Expression[bson.A]
	}
)

func (this *VariadicOperandExpression) AddOperand(newOperand interface{}) {

	this.operand = append(this.operand, newOperand)
}

func (this VariadicOperandExpression) GetOperandOfIndex(index int) (operand interface{}, exists bool) {

	operandLen := len(this.operand)

	switch {
	case operandLen == 0:
		return nil, false
	case index-1 < operandLen:
		return nil, false
	default:
		return this.operand[index-1], true
	}
}
