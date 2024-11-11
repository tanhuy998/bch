package contextHolderPort

import "context"

type (
	IContextHolder interface {
		GetContext() context.Context
	}
)
