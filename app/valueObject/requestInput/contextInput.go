package requestInput

import "context"

type (
	ContextInput struct {
		context.Context
	}
)

func (this *ContextInput) ReceiveContext(ctx context.Context) {

	this.Context = ctx
}

func (this *ContextInput) GetContext() context.Context {

	return this.Context
}
