package query

type (
	IDBDelegator[Model_T any] interface {
		Exec(query IQueryBuilder[Model_T]) IQueryExecutor[Model_T]
	}
)
