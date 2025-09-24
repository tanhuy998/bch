package mongoStorage

import (
	"app/internal/db/storage"
	"context"
)

type (
	aggregate_query_debug_context struct {
		context.Context
		query storage.IArbitraryQuery
	}
)

func resolveDebugContext(ctx context.Context, query storage.IArbitraryQuery) context.Context {

	ret := new(aggregate_query_debug_context)

	ret.Context = ctx
	ret.query = query

	return ret
}

/*
Implement queryDBTracer.IDBQueryTracerDebugContext
*/
func (this *aggregate_query_debug_context) GetDBDebugLog() interface{} {

	if this.query == nil {

		return nil
	}

	return aggregate_debug_log_t{
		QueryType: "aggregate",
		Detail:    this.getDetailQueryDebugLog(),
		Query:     this.query.GetArbitraryQuery(),
	}
}

func (this *aggregate_query_debug_context) getDetailQueryDebugLog() interface{} {

	switch debugger, exist := this.query.(storage.IQueryDebugger); {
	case exist:
		return debugger.GetDetailQueryDebugLog()
	default:
		return nil
	}
}
