package query

import (
	"app/internal/db/storage"
)

type (
	JoinInitFunc = func(queryBuilder IJoinField)

	ISubqueryFilterMethod interface {
		Filter(fn FilterFunc) ISubQueryBuilder //ISubQueryJoinProjectionMethods
	}

	ISubQueryJoinProjectionMethods interface {
		Select(fields ...string) ISubQueryBuilder
	}

	ISubQueryJoinMethod interface {
		Join(another string, fn func(query IJoinField)) ISubQueryBuilder
	}

	IJoinField interface {
		On(localField string, foreignField string) IJoinAlias
	}

	IJoinAlias interface {
		As(name string) ISubQueryBuilder
	}
)

type (
	IJoinMethod[Model_T any] interface {
		//Join(another string, fn func(queryBuilder IJoinField)) IQueryBuilder[Model_T]
		Join(another storage.IDBStorageIdentifier, fn func(queryBuilder IJoinField)) IQueryBuilder[Model_T]
	}

	IJoinTarget[Model_T any] interface {
		On(target string)
	}
)
