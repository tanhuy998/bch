package aggregate

import (
	"app/internal/db/driver/mongoDriver"
	"app/internal/db/query"
	aggregateRelation "app/unitOfWork/aggregate/relation"
	"context"
)

type (
	/*
		AggregateRoot defines base class for aggregates that could do:
		+ managing transaction for database
		+ Query data for a bounded context (which need join between tables/collections of database)
	*/
	AggegateRoot[Aggregate_Mapping_Entity_T any, Source_Repository_Entity_T any] struct {
		AggregateTransaction IAggregateTransaction
		mongoDriver.RelationDelegator[Aggregate_Mapping_Entity_T, Source_Repository_Entity_T]
		default_aggregate_query_builder query.IQueryBuilder[Aggregate_Mapping_Entity_T]
	}
)

func (this *AggegateRoot[Retrieval_Entity_T, Repository_Entity_T]) Query() query.IQueryBuilder[Retrieval_Entity_T] {

	return this.MongoDBDelegator.Query()
}

func (this *AggegateRoot[Retrieval_Entity_T, Repository_Entity_T]) Transaction(
	initCtx context.Context, fn func(context.Context) (interface{}, error),
) (interface{}, error) {

	return this.AggregateTransaction.Transaction(initCtx, fn)
}

/*
Aggregate given relations on first invocation. After that,
*/
func (this *AggegateRoot[Aggregate_Mapping_Entity_T, Source_Repository_Entity_T]) AggregateRelationsOnceOrDefault(
	relations ...aggregateRelation.AbstractRelation,
) query.IQueryBuilder[Aggregate_Mapping_Entity_T] {

	if this.default_aggregate_query_builder == nil {

		this.default_aggregate_query_builder = this.AggregateRelations(
			relations...,
		)
	}

	return this.default_aggregate_query_builder.Clone()
}

/*
Immediately combine database relations and return query builder
*/
func (this *AggegateRoot[Aggregate_Mapping_Entity_T, Source_Repository_Entity_T]) AggregateRelations(
	relations ...aggregateRelation.AbstractRelation,
) query.IQueryBuilder[Aggregate_Mapping_Entity_T] {

	return this.RelationDelegator.AggregateRelations(relations...)
}
