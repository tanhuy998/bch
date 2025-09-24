package logical

import (
	"app/internal/db/driver/mongoDriver/expression"
)

type (
	NotExpression[Expression_T any] struct {
		expression.UnaryExpression[Expression_T]
	}
)

func NewNotLogicalExpression[Expr_T any](
	negatedExpr Expr_T,
) *NotExpression[Expr_T] {

	ret := new(NotExpression[Expr_T])

	ret.SetOperand(negatedExpr)

	return ret
}

func (this NotExpression[Expression_T]) MarshalBSON() ([]byte, error) {

	this.SetOperator("$not")

	return this.UnaryExpression.MarshalBSON()
}
