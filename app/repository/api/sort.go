package repositoryAPI

type (
	ISortInitializer interface {
		Field(name string) ISortOperator
	}

	ISortOperator interface {
		Ascending()
		Descending()
	}

	SortFunc = func(sorter ISortInitializer)

	ISortMethods interface {
		Sort(fn SortFunc)
	}
)
