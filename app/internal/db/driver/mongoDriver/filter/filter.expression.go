package filter

import (
	comparision "app/internal/db/driver/mongoDriver/filter/comparsion"
	"app/internal/db/driver/mongoDriver/filter/date"
	"app/internal/db/driver/mongoDriver/filter/numeric"
	"app/internal/db/driver/mongoDriver/filter/text"
	"app/internal/db/query"
	libCommon "app/internal/lib/common"
)

type (
	IExpressionList interface {
		libCommon.ISliceManipulator[interface{}]
	}
)

type (
	filter_expression_t struct {
		comparision.ComparisionExpression
	}
)

func NewFilterExpression(parent *filter_t, fieldName string) *filter_expression_t {

	if parent == nil {

		panic("parent of casting expression must not be nil")
	}

	ret := new(filter_expression_t)

	ret.SetRoot(
		&filter_expression_pivot_t{
			parent:    parent,
			fieldName: fieldName,
		},
	)

	return ret
}

func (this *filter_expression_t) AsText() query.IFilterTextOperator {

	ret := text.NewFilterFieldTextCasting(
		&this.ComparisionExpression.SecondaryOperator,
	)

	return ret
}

func (this *filter_expression_t) AsNumeric() query.IFilterNumericalOperator {

	ret := numeric.NewFilterExpressionNumericCasting(
		&this.ComparisionExpression.SecondaryOperator,
	)

	return ret
}

func (this *filter_expression_t) AsDate() query.IFilterDateOperator {

	ret := date.NewFilterExpressionDateCasting(
		&this.SecondaryOperator,
	)

	return ret
}

func (this filter_expression_t) MarshalBSON() ([]byte, error) {

	return this.ComparisionExpression.MarshalBSON()
}
