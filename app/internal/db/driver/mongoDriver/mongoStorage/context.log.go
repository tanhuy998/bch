package mongoStorage

import (
	"context"
)

type (
	collection_filterable_operation_debug_log_context_t struct {
		context.Context
		log *collection_filterable_operation_log_t
	}
)

func NewCollectionOperationDebugLogContext(
	//log *collection_filterable_operation_log_t,
	filter interface{},
	ctx context.Context,
) *collection_filterable_operation_debug_log_context_t {

	if ctx == nil {

		ctx = context.TODO()
	}

	return &collection_filterable_operation_debug_log_context_t{
		Context: ctx,
		log:     NewfilterableOpLogLine(filter),
	}
}

func (this *collection_filterable_operation_debug_log_context_t) GetDBDebugLog() interface{} {

	//this.log.AppliedFilter.resolveLog()

	return this.log
}
