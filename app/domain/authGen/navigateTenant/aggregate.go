package navigateTenantDomain

import (
	"app/domain/authGen/navigateTenant/relations"
	"app/model"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/api/crud"

	"github.com/google/uuid"
)

type (
	domain_aggregate struct {
		aggregate.DomainAggregateRoot[model.Tenant, model.Tenant]
		relations.Tenant_TenanAgent_Relation
		relations.Tenant_User_Relation
	}
)

func (this *domain_aggregate) MergeRelations(
	userUUID uuid.UUID,
) crud.IAggregateReader[model.Tenant] {

	return this.ByRelations(
		this.Tenant_TenanAgent_Relation.DispatchDefault(userUUID),
		this.Tenant_User_Relation.DispatchDefault(userUUID),
	)
}
