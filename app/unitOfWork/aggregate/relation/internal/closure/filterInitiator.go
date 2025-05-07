package closure

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
	"app/unitOfWork/aggregate/relation/internal"
	"app/unitOfWork/aggregate/relation/internal/api"
)

type (
	FilterInitiator[
		Resolver_Kind_T api.IResolver,
		Initializer_Kind_T api.IInitializer,
	] struct {
		internal.Initiator[Resolver_Kind_T, Initializer_Kind_T]
		filter_func query.FilterFunc
	}
)

func (this FilterInitiator[Resolver_Kind_T, Initializer_Kind_T]) ResolveRelation(
	local relation.IRelationLocalNavigator, foreign relation.IRelationForeignNavigator,
) {

	this._applyClosureFilter(foreign)

	(this.Initiator.Resolver).ResolveRelation(local, foreign)
}

func (this *FilterInitiator[Resolver_Kind_T, Initializer_Kind_T]) _applyClosureFilter(
	foreignNavigator relation.IRelationForeignNavigator,
) {

	if this.filter_func == nil {

		return
	}

	foreignNavigator.Manipulate(
		func(foreign relation.IReadRelationQueryBuilder) {

			foreign.Filter(this.filter_func)
		},
	)
}

func NewFilterInitiator[Initiator_T relation.IDBRelationInitiator](
	relationInitiator Initiator_T, fn query.FilterFunc,
) *FilterInitiator[relation.IDBRelationInitiator, relation.IDBRelationInitiator] {

	ret := &FilterInitiator[
		relation.IDBRelationInitiator, relation.IDBRelationInitiator,
	]{
		Initiator:   *internal.NewInitiator(relationInitiator),
		filter_func: fn,
	}

	return ret
}

func WithQueryBuilderInterceptor(
	relationInitiator relation.IDBRelationInitiator,
	beforeJoin crud.QueryBuilderFunc,
	afterJoin crud.QueryBuilderFunc,
) relation.IDBRelationInitiator {

	switch {
	case beforeJoin != nil || afterJoin != nil:
		return &interceptor_initiator{
			relationInitiator, beforeJoin, afterJoin,
		}
	default:
		panic("closure.WithQueryBuilderInterceptor's 2nd and 3rd paramemter must not both nil")
	}
}
