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
		Field(name string) IFilterExpressionOperator
	}

	IFilterLogicalOperator interface {
		Or(FilterLogicalGroupFunc)
		And(FilterLogicalGroupFunc)
	}

	IFilterExpressionOperator interface {
		IComaparisonOperator
		Not() IComaparisonOperator
	}

	IComaparisonOperator interface {
		//IFilterLogicalOperator
		Equal(val interface{})
		GreaterThan(val interface{})
		GreaterOrEqual(val interface{})
		LessThan(val interface{})
		LessThanOrEqual(val interface{})
	}

	FilterLogicalGroupFunc = func(filteredField IFilterExpressionOperator)
	FilterFunc             = func(filter IFilterExpression)

	IFilterMethods[Model_T any] interface {
		Filter(FilterFunc) IQueryBuilder[Model_T]
	}
)
