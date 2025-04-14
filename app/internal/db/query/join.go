package query

import "app/internal/db/storage"

type (
	JoinInitFunc func(queryBuilder IJoinField)

	ISubqueryFilterMethod interface {
		Filter(fn FilterFunc) IQueryBuilder //ISubQueryJoinProjectionMethods
	}

	ISubQueryJoinProjectionMethods interface {
		Select(fields ...string) IQueryBuilder
	}

	ISubQueryJoinMethod interface {
		Join(another storage.IDBStorageIdentifier, fn func(query IJoinField)) IJoinUnwindableQueryBuilder // IQueryBuilder
	}

	IJoinField interface {
		On(localField string, foreignField string) IJoinAlias
	}

	IJoinAlias interface {
		As(name string) IQueryBuilder
	}
)

type (
	IJoinMethod[Model_T any] interface {
		//Join(another string, fn func(queryBuilder IJoinField)) IQueryBuilder[Model_T]
		Join(another storage.IDBStorageIdentifier, fn func(queryBuilder IJoinField)) IGenericQueryBuilder[Model_T]
	}

	IJoinTarget[Model_T any] interface {
		On(target string)
	}
)
