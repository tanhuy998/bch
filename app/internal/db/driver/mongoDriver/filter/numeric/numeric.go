package numeric

import (
	"app/internal/db/driver/mongoDriver/filter/operator"
	"app/internal/db/query/helper/inner"
)

type (
	filter_numeric_casting_t struct {
		filter_numeric_iterator_t
	}
)

func NewFilterExpressionNumericCasting(
	operator *operator.SecondaryOperator,
) *filter_numeric_casting_t {

	switch {
	case operator == nil:
		panic("operator of filter_expression_numeric_casting_t must not be nil")
	}

	ret := new(filter_numeric_casting_t)

	ret.Iterator.SetRoot(operator)
	ret.Iterator.GetIterationChain().SetInitialOperand(
		inner.Self(ret.root.Pivot().CurrentFieldName()).GetQuerySelfReference(),
	)

	ret.init()

	return ret
}
