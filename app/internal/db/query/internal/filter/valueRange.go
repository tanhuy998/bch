package filter

import (
	"app/internal/db/query"
	"app/internal/lib/validate"
)

type (
	ValueRangeFilter[T any] struct {
		validator validate.IValueRangeValidator[T]
	}
)

func NewRangeFilter[T any](validator validate.IValueRangeValidator[T]) query.IValueRangeFilter {

	ret := &ValueRangeFilter[T]{
		validator: validator,
	}

	return ret
}

func (this *ValueRangeFilter[T]) ApplyValueRange(p0 query.IFilterExpressionOperator) {

}

func (this *ValueRangeFilter[T]) resolveInbound() {

}

func (this *ValueRangeFilter[T]) resolveOutBound() {

}
