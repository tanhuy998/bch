package query

import "app/internal/db/storage"

type (
	IQueryBuilder interface {
		ISubqueryFilterMethod
		ISubQueryJoinMethod
		ISubQueryProjector
		ISubQueryDataTransform
		ISubQueryDataLimit
		ISkipQueryBuilder
		ISubQueryDataSortOrder
		storage.IArbitraryQuery
		storage.IQueryEndingPhase
	}

	IClonableQueryBuilder interface {
		IQueryBuilder
		Clone() IClonableQueryBuilder
	}
)

type (
	IReadQueryBuilderGenerator interface {
		NewQueryBuilder() IClonableQueryBuilder
	}
)

type (
	IGenericQueryBuilder[Model_T any] interface {
		IJoinMethod[Model_T]
		IFilterMethods[Model_T]
		IProjector[Model_T]
		IClonableQueryExecutor[Model_T]
		IDataTransform[Model_T]
		IDataSortOrder[Model_T]
		Clone() IGenericQueryBuilder[Model_T]
		//IFilterableOperator[Model_T]
	}
)
