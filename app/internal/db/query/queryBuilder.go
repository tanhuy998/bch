package query

type (
	ISubQueryBuilder interface {
		ISubqueryFilterMethod
		ISubQueryJoinMethod
		ISubQueryProjector
		ISubQueryDataTransform
	}
)

type (
	IQueryBuilder[Model_T any] interface {
		IJoinMethod[Model_T]
		IFilterMethods[Model_T]
		IProjector[Model_T]
		IClonableQueryExecutor[Model_T]
		IDataTransform[Model_T]
		IDataSortOrder[Model_T]
		//IFilterableOperator[Model_T]
	}
)
