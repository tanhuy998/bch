package mongoStorage

import "go.mongodb.org/mongo-driver/bson"

type (
	benchmark_log_context_t struct {
		collection_filterable_operation_debug_log_context_t
		bench bson.M
	}
)
