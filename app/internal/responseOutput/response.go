package responseOutput

import "context"

type (
	IResponseMessage interface {
		GetMessage() string
	}

	IResponseContext interface {
		GetContext() context.Context
		SetContext(ctx context.Context)
	}

	ILoggableResponse interface {
		IResponseContext
		IResponseMessage
	}
)
