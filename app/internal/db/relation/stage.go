package relation

import "app/internal/db/query"

type (
	IDBRelationBeforeJoinInterceptor interface {
		InterceptBeforeJoin(queryBuilder query.IQueryBuilder)
	}

	IDBRelationAfterJoinInterceptor interface {
		InterceptAfterJoin(queryBuilder query.IQueryBuilder)
	}
)
