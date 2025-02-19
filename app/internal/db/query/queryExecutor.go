package query

import "context"

type (
	IClonableQueryExecutor[Model_T any] interface {
		NewExecutor(dbDelegator IDBDelegator[Model_T]) IQueryExecutor[Model_T]
		IQueryExecutor[Model_T]
	}

	IQueryExecutor[Model_T any] interface {
		IQueryPaginate[Model_T]
		First(ctx context.Context) (*Model_T, error)
		All(ctx context.Context) ([]Model_T, error)
	}
)
