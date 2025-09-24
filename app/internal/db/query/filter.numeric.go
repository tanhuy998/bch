package query

type (
	IFilterArithmetricOperator interface {
		Add(val interface{}) IFilterNumericalComparisonOperator
		Substract(val interface{}) IFilterNumericalComparisonOperator
		Multiply(val interface{}) IFilterNumericalComparisonOperator
		Divide(val interface{}) IFilterNumericalComparisonOperator
	}

	IFilterNumericalComparisonOperator interface {
		IGenericNegationOperator[IComparisonOperator]
		IComparisonOperator
		IFilterArithmetricOperator
	}

	IFilterNumericalOperator IFilterNumericalComparisonOperator
)
