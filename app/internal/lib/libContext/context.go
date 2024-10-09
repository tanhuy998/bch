package libContext

import (
	"context"
)

type (
	NoReadContext struct {
		context.Context
	}
)

func WrapNoReadContext(ctx context.Context) *NoReadContext {

	return &NoReadContext{ctx}
}

func IsNoRead(ctx context.Context) bool {

	_, ok := ctx.(*NoReadContext)

	return ok
}
