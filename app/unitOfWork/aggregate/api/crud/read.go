package crud

import "app/internal/db/query"

type (
	QueryBuilderFunc func(queryBuilder query.IQueryBuilder)
)

type (
	IAggregateReader[Read_Entity_T any] interface {
		Read(
			initFn QueryBuilderFunc,
		) query.IGenericQueryExecutor[Read_Entity_T]
	}
)
