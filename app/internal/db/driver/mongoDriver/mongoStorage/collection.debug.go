package mongoStorage

import (
	dbQueryTracerPort "app/port/dbQueryTracer"
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	debug_collection_t struct {
		AbstractMongoCollection
		//collection *mongo.Collection
		Tracer dbQueryTracerPort.IDBQueryTracer
	}
)

func (this *debug_collection_t) SetTracer(t dbQueryTracerPort.IDBQueryTracer) {

	if t == nil {

		panic("debug_collection_t error: tracer must not be nil")
	}

	this.Tracer = t
}

func (this *debug_collection_t) prepareOpFilterDebugLogContext(
	filter interface{}, ctx context.Context,
) *collection_filterable_operation_debug_log_context_t {

	return NewCollectionOperationDebugLogContext(
		// &collection_filterable_operation_log_t{
		// 	AppliedFilter: filter_debug_log_t{
		// 		JsonForm: filter,
		// 	},
		// },
		filter,
		ctx,
	)
}

func (this *debug_collection_t) prepareDebugOptions(opts interface{}, ctx context.Context) {

	switch opts.(type) {
	case *options.AggregateOptions:
	case *options.FindOptions:
	case *options.UpdateOptions:
	case *options.DeleteOptions:
	case *options.FindOneOptions:
	default:
		return
	}
}
