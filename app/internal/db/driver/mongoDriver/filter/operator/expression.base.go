package operator

import (
	"app/internal/db/driver/mongoDriver/expression"
	"app/internal/db/query"
)

const (
	lhs_operand_index = 0
	rhs_operand_index = 1
)

type (
	base_t struct {
		expression.Expression[interface{}]
		operands [2]interface{}
	}
)

func (this *base_t) Init() {

	this.initOperands()
}

func (this *base_t) initOperands() {

	if this.Expression.GetOperand() != nil {

		return
	}

	this.Expression.SetOperand(this.operands[:])
}

func (this *base_t) SetOperator(operator string) {

	this.Expression.SetOperator(operator)
}

func (this base_t) GetOperator() string {

	return this.Expression.GetOperator()
}

func (this *base_t) SetRightOperand(rhs interface{}) {

	this.initOperands()

	switch v := rhs.(type) {
	case query.IQuerySelfReference:
		this.operands[rhs_operand_index] = v.GetQuerySelfReference()
	default:
		this.operands[rhs_operand_index] = v
	}
}

func (this *base_t) SetLeftOperand(lhs interface{}) {

	this.initOperands()

	switch v := lhs.(type) {
	case query.IQuerySelfReference:
		this.operands[lhs_operand_index] = v.GetQuerySelfReference()
	default:
		this.operands[lhs_operand_index] = v
	}
}

func (this base_t) GetLeftOperand() interface{} {

	this.initOperands()

	return this.operands[lhs_operand_index]
}

func (this base_t) GetRightOperand() interface{} {

	this.initOperands()

	return this.operands[rhs_operand_index]
}
