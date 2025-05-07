package closure

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
)

type (
	interceptor_initiator struct {
		relation.IDBRelationInitiator
		intercept_before_fn crud.QueryBuilderFunc
		intercept_after_fn  crud.QueryBuilderFunc
	}
)

func (this interceptor_initiator) InterceptBeforeJoin(queryBuilder query.IQueryBuilder) {

	switch this.intercept_before_fn {
	case nil:
		return
	default:
		this.intercept_before_fn(queryBuilder)
	}
}

func (this interceptor_initiator) InterceptAfterJoin(queryBuilder query.IQueryBuilder) {

	switch this.intercept_after_fn {
	case nil:
		return
	default:
		this.intercept_after_fn(queryBuilder)
	}
}
