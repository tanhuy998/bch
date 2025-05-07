package getUserParticipatedCommandGroupsDomain

import (
	"app/model"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/api/crud"

	"github.com/google/uuid"
)

type (
	domain_aggregate struct {
		aggregate.DomainAggregateRoot[model.CommandGroup, model.CommandGroup]
		CommandGroup_CommandGroupUser_Relation
	}
)

func (this *domain_aggregate) MergeRelations(tenantUUID uuid.UUID, userUUID uuid.UUID) crud.IAggregateReader[model.CommandGroup] {

	return this.ByRelations(
		this.CommandGroup_CommandGroupUser_Relation.DispatchDefault(tenantUUID, userUUID),
	)
}
