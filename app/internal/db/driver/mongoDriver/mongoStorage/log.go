package mongoStorage

// type (
// 	IMongoDebugLoggableContext interface {
// 		GetMongoDebogLog() interface{}
// 	}
// )

type (
	aggregate_debug_log struct {
		QueryType string      `json:"query_type"`
		Detail    interface{} `json:"detail,omitempty"`
		Query     interface{} `json:"query,omitempty"`
	}
)
