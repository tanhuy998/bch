package closure

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
	"app/unitOfWork/aggregate/relation/internal/closure"
)

func WithForeignFilter(
	relationInitiator relation.IDBRelationInitiator,
	fn query.FilterFunc,
) relation.IDBRelationInitiator {

	return closure.NewFilterInitiator(relationInitiator, fn)
}

// func WithClosure(
// 	relationInitiator relation.IDBRelationInitiator,
// )

func WithQueryBuilderInterceptor(
	relationInitiator relation.IDBRelationInitiator,
	beforeJoin crud.QueryBuilderFunc,
	afterJoin crud.QueryBuilderFunc,
) relation.IDBRelationInitiator {

	return closure.WithQueryBuilderInterceptor(
		relationInitiator, beforeJoin, afterJoin,
	)
}
