package aggregate

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/api/crud"
	"app/unitOfWork/aggregate/relation/lib/closure"
)

type (
	DomainAggregateRoot[Read_Entity_T any, Local_Storage_Unit_Entity_T any] struct {
		AggegateRoot[Read_Entity_T, Local_Storage_Unit_Entity_T]
	}
)

func (this *DomainAggregateRoot[Read_Entity_T, Local_Storage_Unit_Entity_T]) MergeRelationsByDomain(
	ctx IDomainContext, relationIntitiators ...relation.IDBRelationInitiator,
) crud.IAggregateReader[Read_Entity_T] {

	var filter_func query.FilterFunc = func(filter query.IFilterExpression) {

		filter.Field("tenantUUID").Equal(ctx.GetTenantUUID())
	}

	transformedIntitiator := make([]relation.IDBRelationInitiator, len(relationIntitiators))

	for i, oldInitiator := range relationIntitiators {

		transformedIntitiator[i] = closure.WithForeignFilter(
			oldInitiator, filter_func,
		)
	}

	return this.ByRelations(transformedIntitiator...)
}
