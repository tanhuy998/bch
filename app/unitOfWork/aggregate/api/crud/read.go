package crud

import "app/internal/db/query"

type (
	IAggregateReader[Read_Entity_T any] interface {
		Read(
			initFn func(queryBuilder query.IQueryBuilder),
		) query.IGenericQueryExecutor[Read_Entity_T]
	}
)
