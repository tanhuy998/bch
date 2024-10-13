package domain

import "context"

type (
	ContextualDomainService[Domain_Context_T context.Context] struct {
	}
)

func (this ContextualDomainService[Domain_Context_T]) InDomainContext(ctx context.Context) bool {

	_, ok := any(ctx).(*Domain_Context_T)

	return ok
}
