package valueRange

import (
	"app/internal/db/query"
	"app/internal/db/query/internal/filter"
	"app/internal/lib/validate"
	"cmp"
)

func ForComparableType[T validate.IComparable[T]](
	lowerBound *T, upperBound *T,
) query.IValueRangeFilter {

	return filter.NewRangeFilter(
		validate.NewComparableTypeRangePaginator(lowerBound, upperBound),
	)
}

func ForPrimitiveType[T cmp.Ordered](
	lowerBound *T, upperBound *T,
) query.IValueRangeFilter {

	return filter.NewRangeFilter(
		validate.NewPrimitiveTypeDataRangePaginator(lowerBound, upperBound),
	)
}
