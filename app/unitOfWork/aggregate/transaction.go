package aggregate

import "context"

type (
	IAggregateTransaction interface {
		Transaction(initCtx context.Context, context func(tranCtx context.Context) (interface{}, error)) (interface{}, error)
	}
)
