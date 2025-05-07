package query

type (
	IGenericNegationOperator[Negatable_T any] interface {
		Not() Negatable_T
	}
)

type (
	/*
		FILTER EXPRESSION DEFINITIONS
	*/

	IDataConditionFilterExpression interface {
		//IConditionLogicalOperator
		Field(string) INegatableDataConditionComparisonOperator
	}

	IConditionExpression interface {
		IConditionLogicalOperator
		IDataConditionFilterExpression
	}

	//ConditionLogicalOperatorFunc func(where )

)

type (
	/*
		COMPARISON OPERATOR DEFINITIONS
	*/

	IDataConditionComparisonOperator interface {
		//Not() IDataConditionComparisonOperator
		//IGenericConditionNotOperator[IDataConditionComparisonOperator]
		IComaparisonOperator
		IComparisonRange
		INumericalRangeComparisonOperator
		//IConditionLogicalOperator
	}

	INegatableDataConditionComparisonOperator interface {
		IDataConditionComparisonOperator
		IGenericNegationOperator[IDataConditionComparisonOperator]
	}
)

type (
	/*
		RANGE COMPARISON OPERATOR DEFINITIONS
	*/

	IComparisonRange interface {
		EqualOneOfVals(vals ...interface{})
	}

	INegatableComparisonRange interface {
		IComparisonRange
		IGenericNegationOperator[IComparisonRange]
	}

	INumericalRangeComparisonOperator interface {
		InNumericalIntegerRange(minVal int64, maxVal int64)
		InNumericalUnsignedIntergerRange(min uint64, max uint64)
		InNumericalFloatingPointRange(min float64, max float64)
	}

	INegatableNumericalRangeComparisonOperator interface {
		IGenericNegationOperator[INumericalRangeComparisonOperator]
		INumericalRangeComparisonOperator
	}
)

type (
	/*
		LOGICAL OPERATOR DEFINITIONS
	*/

	IConditionLogicalOperator interface {
		//Not() IConditionLogicalOperator
		And(conditionExpressions ...DataConditionMatchFunc) IDataConditionExpressionResult
		Or(conditionExpressions ...DataConditionMatchFunc) IDataConditionExpressionResult
	}

	INegatableLogicalOperator interface {
		IConditionLogicalOperator
		IGenericNegationOperator[IConditionLogicalOperator]
	}
)

type (
	IDataConditionExpressionResult interface {
		GetCondtionExpression() interface{}
	}

	DataConditionMatchFunc func(expression IGeneralDataConditionExpression) IDataConditionExpressionResult

	LogicalExpressionInitFunc func(logical INegatableLogicalOperator) IDataConditionExpressionResult

	FilterExpressInitFunc func(filter IDataConditionFilterExpression)

	IGeneralDataConditionExpression interface {
		//AsLogical(fn LogicalExpressionInitFunc) IDataConditionExpressionResult
		Logical() INegatableLogicalOperator
		Filter(fn FilterExpressInitFunc) IDataConditionExpressionResult
		// INegatableLogicalOperator
		// IDataConditionFilterExpression
	}

	IDataCondition interface {
		Match(fn DataConditionMatchFunc) IQueryBuilder
	}
)

// func t() {

// 	var a IDataCondition

// 	a.Match(
// 		func(expression IGeneralDataConditionExpression) IDataConditionExpressionResult {

// 			return expression.Filter(
// 				func(filter IDataConditionFilterExpression) {

// 				},
// 			)
// 		},
// 	)
// }
