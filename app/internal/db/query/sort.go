package query

type (
	ISortInitializer interface {
		Field(name string) ISortOperator
	}

	ISortOperator interface {
		Ascending()
		Descending()
	}

	SortFunc = func(sorter ISortInitializer)

	ISubQueryDataSortOrder interface {
		SortOrder(fn SortFunc) IQueryBuilder
	}
)

type (
	IDataSortOrder[Model_T any] interface {
		SortOrder(fn SortFunc) IGenericQueryBuilder[Model_T]
	}
)
