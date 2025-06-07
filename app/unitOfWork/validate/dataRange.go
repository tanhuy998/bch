package validate

import "cmp"

type (
	DataRangeValidator[Data_T cmp.Ordered] struct {
		lower_bound *Data_T
		upper_bound *Data_T
	}
)

func NewDataRangeValidator[Data_T cmp.Ordered](lowerBound *Data_T, upperBound *Data_T) *DataRangeValidator[Data_T] {

	return &DataRangeValidator[Data_T]{
		lower_bound: lowerBound,
		upper_bound: upperBound,
	}
}

func (this DataRangeValidator[Data_T]) IsInboundRange() bool {

	switch {
	case this.lower_bound == nil:
		return false
	case this.upper_bound == nil:
		return false
	case *this.lower_bound < *this.upper_bound:
		return true
	default:
		return false
	}
}

func (this DataRangeValidator[Data_T]) IsOutboundRange() bool {

	switch {
	case this.lower_bound == nil:
		return false
	case this.upper_bound == nil:
		return false
	case *this.lower_bound > *this.upper_bound:
		return true
	default:
		return false
	}
}

func (this DataRangeValidator[Data_T]) IsUpperboundRange() bool {

	return this.upper_bound != nil && this.lower_bound == nil
}

func (this DataRangeValidator[Data_T]) IsLowerboundRange() bool {

	return this.lower_bound != nil && this.upper_bound == nil
}

func (this DataRangeValidator[Data_T]) GetLowerBound() Data_T {

	return *this.lower_bound
}

func (this DataRangeValidator[Data_T]) GetUpperBound() Data_T {

	return *this.upper_bound
}
