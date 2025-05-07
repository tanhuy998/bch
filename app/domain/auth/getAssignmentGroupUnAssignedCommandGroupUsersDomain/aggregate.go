package getAssignmentGroupUnAssignedCommandGroupUsersDomain

import (
	"app/domain/auth/getAssignmentGroupUnAssignedCommandGroupUsersDomain/relations"
	"app/model"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/api/crud"

	"github.com/google/uuid"
)

type (
	Domain_Aggregate struct {
		aggregate.DomainAggregateRoot[model.CommandGroupUser, model.CommandGroupUser]
		relations.CommandGroupUser_AssignmentGroup_Relation
		relations.CommandGroupUser_Participated_AssignmentGroup_Relation
		relations.CommandGroupUser_User_Relation
		relations.CommandGroupUser_CreatedUser_Relation
	}
)

func (this *Domain_Aggregate) MergeDefaultRelation(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, excludedCommandGroupUserUUIDs []uuid.UUID,
) crud.IAggregateReader[model.CommandGroupUser] {

	return this.ByRelations(
		this.CommandGroupUser_AssignmentGroup_Relation.DispatchDefault(tenantUUID, assignmentUUID, excludedCommandGroupUserUUIDs),
		this.CommandGroupUser_Participated_AssignmentGroup_Relation.DispatchDefault(tenantUUID, assignmentUUID),
		this.CommandGroupUser_User_Relation.DispatchDefault(tenantUUID),
		this.CommandGroupUser_CreatedUser_Relation.DispatchDefault(tenantUUID),
	)
}

func (this *Domain_Aggregate) MergeRelationsForLookupUnAssignedCommandGroupUsers(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID, lookupCommandGroupUserUUIDs []uuid.UUID,
) crud.IAggregateReader[model.CommandGroupUser] {

	return this.ByRelations(
		this.CommandGroupUser_AssignmentGroup_Relation.DispatchForLookupUnAssigned(tenantUUID, assignmentUUID, lookupCommandGroupUserUUIDs),
		this.CommandGroupUser_Participated_AssignmentGroup_Relation.DispatchForLookupUnAssigned(tenantUUID, assignmentUUID),
		this.CommandGroupUser_User_Relation.DispatchForLookupUnAssigned(tenantUUID),
		this.CommandGroupUser_CreatedUser_Relation.DispatchForLookupUnAssigned(tenantUUID),
	)
}
