package query

type (
	IComparable[T any] interface {
		Compare(T) int
	}

	IComparableDataRange[T IComparable[T]] interface {
		GetLowerBound() interface{}
		GetUpperBound() interface{}
	}

	IComparableDataRangeValidator interface {
		IsInboundRange() bool
		IsOutboundRange() bool
		IsUpperboundRange() bool
		IsLowerboundRange() bool
	}
)
