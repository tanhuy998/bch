package closure

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal"
	"app/unitOfWork/aggregate/relation/internal/api"
)

type (
	ClosureGenerator[
		Resolver_Kind_T api.IResolver,
		Initializer_Kind_T api.IInitializer,
	] struct {
		internal.Initiator[Resolver_Kind_T, Initializer_Kind_T]
	}
)

func (this *ClosureGenerator[Resolver_Kind_T, Initializer_Kind_T]) WithForeignFilter(
	fn query.FilterFunc,
) relation.IDBRelationInitiator {

	if fn == nil {

		panic("closure filter coudld not be nil")
	}

	ret := FilterInitiator[Resolver_Kind_T, Initializer_Kind_T]{
		Initiator:   this.Initiator,
		filter_func: fn,
	}

	return ret
}
