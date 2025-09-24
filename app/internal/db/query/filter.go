package query

type ()

type (
	// IFilterRepository[Model_T any] interface {
	// 	Filter(filter interface{}) ICRUDRepository[Model_T]
	// }

	IFilterableQueryBuilder[Model_T any] interface {
	}

	IFilterableOperator[Model_T any] interface {
		//IRepositoryReadOperator[Model_T]
		//IRepositoryProjectableOperator[Model_T]
		IProjector[Model_T]
		IDataRetrievalQuery[Model_T]
		//IDataMutationQuery[Model_T]
		//IDataRemovalQuery[Model_T]
	}

	IFilterGenerator interface {
		//Add(...interface{}) IFilterGenerator
		IFilterExpression
	}

	IFilterExpression interface {
		Field(name string) INegatableDataConditionComparisonOperator //IFilterExpressionOperator
	}

	IFilterLogicalOperator interface {
		Or(FilterLogicalGroupFunc)
		And(FilterLogicalGroupFunc)
	}

	IFilterExpressionOperator interface {
		IComparisonOperator
		Not() IComparisonOperator
	}

	// IFilterExpressionOperator internalQuery.IFilterExpressionOperator

	IComparisonOperator interface {
		IComparisonRange
		//IFilterLogicalOperator
		Equal(val interface{})
		GreaterThan(val interface{})
		GreaterOrEqual(val interface{})
		LessThan(val interface{})
		LessThanOrEqual(val interface{})
		In(vals ...interface{})
	}

	IFilterFieldCastedType interface {
		AsText() IFilterTextOperator
		AsNumeric() IFilterNumericalOperator
		AsDate() IFilterDateOperator
	}

	// IComaparisonOperator internalQuery.IComaparisonOperator

	IComparisonRangeOperator interface {
		ByRange(IValueRangeFilter)
	}

	FilterLogicalGroupFunc func(filteredField IFilterExpressionOperator)
	FilterFunc             func(filter IFilterExpression)

	IFilterMethods[Model_T any] interface {
		Filter(FilterFunc) IGenericQueryBuilder[Model_T]
	}
)

type (
	ITextSearchOperator interface {
		Like(str string)
	}
)

type (
	IValueRangeFilter interface {
		ApplyValueRange(IFilterExpressionOperator)
	}
)
