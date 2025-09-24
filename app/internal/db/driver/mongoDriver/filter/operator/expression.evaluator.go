package operator

import "app/internal/db/driver/mongoDriver/expression"

type (
	expression_evaluator_t struct {
		base_t
	}
)

func (this *expression_evaluator_t) AsBinaryExpression() expression.IBinaryExpression[interface{}, interface{}] {

	return this
}

func (this *expression_evaluator_t) AsUnaryExpression() expression.IUnaryExpression[interface{}] {

	return &this.Expression
}
