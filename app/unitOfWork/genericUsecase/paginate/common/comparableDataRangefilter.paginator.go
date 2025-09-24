package common

import (
	"app/internal/db/query"
	"app/unitOfWork/validate"
)

type (
	ComparableDataRangePaginatorFilterResolver[T validate.IComparable[T]] struct {
		validate.ComparableDataRangeValidator[T]
		dataTransformFn func(*T) *T
		filterField     string
	}
)

func NewComparableDataRangePaginator[T validate.IComparable[T]](
	filterField string, lowerBoundData *T, upperBoundData *T,
) *ComparableDataRangePaginatorFilterResolver[T] {

	ret := &ComparableDataRangePaginatorFilterResolver[T]{
		filterField: filterField,
	}

	ret.ComparableDataRangeValidator = *validate.NewComparableDataRangePaginator(lowerBoundData, upperBoundData)

	return ret
}

func (this *ComparableDataRangePaginatorFilterResolver[T]) SetFilterField(name string) {

	this.filterField = name
}

func (this *ComparableDataRangePaginatorFilterResolver[T]) SetDataTransformFunc(
	fn func(*T) *T,
) {

	this.dataTransformFn = fn
}

func (this *ComparableDataRangePaginatorFilterResolver[T]) ResolveDataRangeFilter(
	initExpression query.IGeneralDataConditionExpression,
) query.IDataConditionExpressionResult {

	switch {
	case this.IsOutboundRange():
		return this.resolveOutBoundFilter(initExpression)
	default:
		return this.resolveGeneralFilter(initExpression)
	}
}

func (this *ComparableDataRangePaginatorFilterResolver[T]) resolveOutBoundFilter(
	initExpression query.IGeneralDataConditionExpression,
) query.IDataConditionExpressionResult {

	if !this.IsOutboundRange() {

		return nil
	}

	return initExpression.Logical().Or(
		func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

			// return expression.Filter(
			// 	func(filter query.IDataConditionFilterExpression) {

			// 		filter.Field(this.filterField).GreaterOrEqual(*this.GetLowerBound())
			// 	},
			// )

			switch expectedVal := this.tryTransformComparedData(this.GetLowerBound()); {
			case expectedVal == nil:
				return nil
			default:
				return expression.Filter(
					func(filter query.IDataConditionFilterExpression) {

						filter.Field(this.filterField).GreaterOrEqual(expectedVal)
					},
				)
			}
		},
		func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

			// return expression.Filter(
			// 	func(filter query.IDataConditionFilterExpression) {

			// 		filter.Field(this.filterField).LessThanOrEqual(*this.GetUpperBound())
			// 	},
			// )

			switch expectedVal := this.tryTransformComparedData(this.GetUpperBound()); {
			case expectedVal == nil:
				return nil
			default:
				return expression.Filter(
					func(filter query.IDataConditionFilterExpression) {

						filter.Field(this.filterField).GreaterOrEqual(expectedVal)
					},
				)
			}
		},
	)
}

func (this *ComparableDataRangePaginatorFilterResolver[T]) resolveGeneralFilter(
	initExpression query.IGeneralDataConditionExpression,
) query.IDataConditionExpressionResult {

	if this.IsOutboundRange() {

		return nil
	}

	return initExpression.Logical().And(
		func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

			// var (
			// 	expectedVal *T
			// )

			// lowerBound := this.GetLowerBound()

			// switch {
			// case lowerBound == nil && this.dataTransformFn == nil:
			// 	return nil
			// case this.dataTransformFn != nil:
			// 	expectedVal = this.dataTransformFn(lowerBound)
			// 	fallthrough
			// case expectedVal == nil:
			// 	return nil
			// default:
			// 	return expression.Filter(
			// 		func(filter query.IDataConditionFilterExpression) {

			// 			filter.Field(this.filterField).GreaterOrEqual(expectedVal)
			// 		},
			// 	)
			// }

			switch expectedVal := this.tryTransformComparedData(this.GetLowerBound()); {
			case expectedVal == nil:
				return nil
			default:
				return expression.Filter(
					func(filter query.IDataConditionFilterExpression) {

						filter.Field(this.filterField).GreaterOrEqual(expectedVal)
					},
				)
			}
		},
		func(expression query.IGeneralDataConditionExpression) query.IDataConditionExpressionResult {

			// var (
			// 	expectedVal *T
			// )

			// upperBound := this.GetUpperBound()

			// switch {
			// case upperBound == nil && this.dataTransformFn == nil:
			// 	return nil
			// case this.dataTransformFn != nil:
			// 	expectedVal = this.dataTransformFn(upperBound)
			// 	fallthrough
			// case expectedVal == nil:
			// 	return nil
			// default:
			// 	return expression.Filter(
			// 		func(filter query.IDataConditionFilterExpression) {

			// 			filter.Field(this.filterField).LessThanOrEqual(expectedVal)
			// 		},
			// 	)
			// }

			switch expectedVal := this.tryTransformComparedData(this.GetUpperBound()); {
			case expectedVal == nil:
				return nil
			default:
				return expression.Filter(
					func(filter query.IDataConditionFilterExpression) {

						filter.Field(this.filterField).LessThanOrEqual(expectedVal)
					},
				)
			}
		},
	)
}

func (this *ComparableDataRangePaginatorFilterResolver[T]) tryTransformComparedData(comparedVal *T) (ret *T) {

	switch {
	case this.dataTransformFn == nil:
		return comparedVal
	default:
		return this.dataTransformFn(comparedVal)
	}
}
