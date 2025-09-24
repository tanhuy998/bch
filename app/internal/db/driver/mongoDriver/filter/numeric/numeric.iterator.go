package numeric

import (
	"app/internal/db/driver/mongoDriver/filter/chain"
	"app/internal/db/driver/mongoDriver/filter/iterator"
	"app/internal/db/driver/mongoDriver/filter/numeric/arithmetric"
	"app/internal/db/driver/mongoDriver/filter/operator"
	"app/internal/db/query"
)

type (
	filter_numeric_iterator_t struct {
		numeric_operator_t
		iterator.Iterator[
			*arithmetric.ArithmetricChain,
			chain.ExpressionChainNode,
		]
		arithmetricExpressionChain arithmetric.ArithmetricChain
	}
)

func (this *filter_numeric_iterator_t) init() {

}

func (this *filter_numeric_iterator_t) SetPrimaryOperator(operator *operator.SecondaryOperator) {

	this.numeric_operator_t.root = operator
	this.Iterator.SetRoot(operator)
}

func (this *filter_numeric_iterator_t) Add(val interface{}) query.IFilterNumericalComparisonOperator {

	this.Iterator.GetIterationChain().Add(val)

	return this
}

func (this *filter_numeric_iterator_t) Substract(val interface{}) query.IFilterNumericalComparisonOperator {

	this.Iterator.GetIterationChain().Substract(val)

	return this
}

func (this *filter_numeric_iterator_t) Multiply(val interface{}) query.IFilterNumericalComparisonOperator {

	this.Iterator.GetIterationChain().Multiply(val)

	return this
}

func (this *filter_numeric_iterator_t) Divide(val interface{}) query.IFilterNumericalComparisonOperator {

	this.Iterator.GetIterationChain().Divide(val)

	return this
}

func (this *filter_numeric_iterator_t) MarshalBSON() ([]byte, error) {

	return this.Root().MarshalBSON()
}
