package expression

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	expression_operand_t[Operand_T any] struct {
		operand Operand_T
	}
)

func (this *expression_operand_t[Operand_T]) SetOperand(operand Operand_T) {

	this.operand = operand
}

func (this expression_operand_t[Operand_T]) GetOperand() Operand_T {

	return this.operand
}

func (this expression_operand_t[Operand_T]) MarshalBSON() ([]byte, error) {

	return bson.Marshal(this.operand)
}

func (this expression_operand_t[Operand_T]) MarshalJSON() ([]byte, error) {

	return json.Marshal(this.operand)
}
