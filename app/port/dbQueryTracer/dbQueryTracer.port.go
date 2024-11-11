package dbQueryTracerPort

import "context"

type (
	IDBQueryTracer interface {
		Trace(collectionName string, label string, ctx context.Context) (stop func(error))
	}
)
