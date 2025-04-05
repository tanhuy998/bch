package mongoStorage

import (
	"app/internal/db/driver/mongoDriver/lib"
	"app/internal/db/storage"
	"context"
)

type (
	QueryExecutorProxy[Entity_T any] struct {
		//MongoDBQueryMonitorCollection
		MongoStorageUnit[Entity_T]
	}
)

func (this *QueryExecutorProxy[Entity_T]) ToSlice(
	result interface{}, query storage.IArbitraryQuery, ctx context.Context,
) error {

	ctx = resolveDebugContext(ctx, query)

	query.Done()

	return lib.AggregateRaw(
		result, &this.MongoDBQueryMonitorCollection, query.GetArbitraryQuery(), ctx,
	)
}

func (this *QueryExecutorProxy[Entity_T]) First(
	result interface{}, query storage.IArbitraryQuery, ctx context.Context,
) error {

	ctx = resolveDebugContext(ctx, query)

	query.Done()

	return lib.AggregateRawOne(
		result, &this.MongoDBQueryMonitorCollection, query.GetArbitraryQuery(), ctx,
	)
}
