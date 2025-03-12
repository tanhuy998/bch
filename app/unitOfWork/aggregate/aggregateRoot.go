package aggregate

import (
	"context"
)

type (
	/*
		AggregateRoot defines base class for aggregates that could do:
		+ managing transaction for database
		+ Query data for a bounded context (which need join between tables/collections of database)
	*/
	AggegateRoot[Read_Entity_T any, Local_Storage_Unit_Entity_T any] struct {
		DefaultAggregateRelation[Read_Entity_T, Local_Storage_Unit_Entity_T]
		AggregateTransaction IAggregateTransaction

		//mongoDriver.RelationDelegator[Read_Entity_T, Local_Storage_Unit_Entity_T]
		//storage.IDBStorageQueryExecutor[Read_Entity_T]
		//default_aggregate_query_builder query.IGenericQueryBuilder[Read_Entity_T]
	}
)

// func (this *AggegateRoot[Retrieval_Entity_T, Repository_Entity_T]) Query() query.IGenericQueryBuilder[Retrieval_Entity_T] {

// 	return this.MongoDBDelegator.Query()
// }

func (this *AggegateRoot[Retrieval_Entity_T, Repository_Entity_T]) Transaction(
	initCtx context.Context, fn func(context.Context) (interface{}, error),
) (interface{}, error) {

	return this.AggregateTransaction.Transaction(initCtx, fn)
}

// /*
// Aggregate given relations on first invocation. After that,
// */
// func (this *AggegateRoot[Aggregate_Mapping_Entity_T, Source_Repository_Entity_T]) AggregateRelationsOnceOrDefault(
// 	relations ...aggregateRelation.AbstractRelation,
// ) query.IGenericQueryBuilder[Aggregate_Mapping_Entity_T] {

// 	if this.default_aggregate_query_builder == nil {

// 		this.default_aggregate_query_builder = this.AggregateRelations(
// 			relations...,
// 		)
// 	}

// 	return this.default_aggregate_query_builder.Clone()
// }

// /*
// Immediately combine database relations and return query builder
// */
// func (this *AggegateRoot[Aggregate_Mapping_Entity_T, Source_Repository_Entity_T]) AggregateRelations(
// 	relations ...aggregateRelation.AbstractRelation,
// ) query.IGenericQueryBuilder[Aggregate_Mapping_Entity_T] {

// 	return this.RelationDelegator.AggregateRelations(relations...)
// }
