package query

type (
	IFilterMathOperator interface {
		Add(val interface{}) IFilterNumericalComparisonOperator
		Minus(val interface{}) IFilterExpressionOperator
		Multiply(val interface{}) IFilterExpressionOperator
		Divide(val interface{}) IFilterExpressionOperator
	}

	IFilterNumericalComparisonOperator interface {
		IGenericNegationOperator[IFilterNumericalComparisonOperator]
		Equal(val interface{})
		GreaterThan(val interface{})
		GreaterOrEqual(val interface{})
		LessThan(val interface{})
		LessThanOrEqual(val interface{})
		In(vals ...interface{})
	}

	IFilterNumericalOperator IFilterNumericalComparisonOperator
)
