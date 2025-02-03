package aggregate

import (
	"app/internal/db/driver/mongoDriver"
	"app/internal/db/query"
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
		DBDelegator          mongoDriver.MongoDBDelegator[Aggregate_Mapping_Entity_T, Source_Repository_Entity_T]
	}
)

func (this *AggegateRoot[Retrieval_Entity_T, Repository_Entity_T]) Query() query.IQueryBuilder[Retrieval_Entity_T] {

	return this.DBDelegator.Query()
}

func (this *AggegateRoot[Retrieval_Entity_T, Repository_Entity_T]) Transaction(initCtx context.Context, fn func(context.Context) (interface{}, error)) (interface{}, error) {

	return this.AggregateTransaction.Transaction(initCtx, fn)
}
