package query

type (
	IDBDelegator[Model_T any] interface {
		Exec(query IGenericQueryBuilder[Model_T]) IQueryExecutor[Model_T]
	}
)
