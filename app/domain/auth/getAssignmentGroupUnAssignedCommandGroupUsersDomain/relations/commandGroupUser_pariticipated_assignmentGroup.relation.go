package relations

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/relation/facade/binaryRelationFacade"
	binaryRelationTriat "app/unitOfWork/aggregate/relation/triat/binary"

	"github.com/google/uuid"
)

type (
	CommandGroupUserParticipatedAssignmentGroupContext interface {
		aggregate.IDomainContext
		aggregate.IAssignmentGroupDomain
	}
)

type (
	CommandGroupUser_Participated_AssignmentGroup_Relation struct {
		binaryRelationTriat.OneToManyBy[model.AssignmentGroupMember, model.AssignmentGroup]
		tenantUUID     uuid.UUID
		assignmentUUID uuid.UUID
	}
)

func (this CommandGroupUser_Participated_AssignmentGroup_Relation) GetForeignInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("assignmentGroupUUID", "uuid").As("assignmentGroups").
			Filter(
				func(filter query.IFilterExpression) {
					filter.Field("tenantUUID").Equal(this.tenantUUID)
					filter.Field("assignmentUUID").Not().Equal(this.assignmentUUID)
				},
			)
	}
}

func (this CommandGroupUser_Participated_AssignmentGroup_Relation) GetBridgeInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("uuid", "commandGroupUserUUID").As("assignmentGroupMembers").
			Filter(
				func(filter query.IFilterExpression) {

					filter.Field("tenantUUID").Equal(this.tenantUUID)
				},
			)
	}
}

func (this *CommandGroupUser_Participated_AssignmentGroup_Relation) DispatchDefault(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID,
) relation.IDBRelationInitiator {

	clone := *this

	clone.tenantUUID = tenantUUID
	clone.assignmentUUID = assignmentUUID

	// return binaryRelationFacade.MarshallWithFilter(
	// 	ref,
	// 	func(filter query.IFilterExpression) {

	// 		filter.Field("tenantUUID").Equal(tenantUUID)
	// 	},
	// 	func(filter query.IFilterExpression) {

	// 		filter.Field("tenantUUID").Equal(tenantUUID)
	// 		filter.Field("assignmentUUID").Not().Equal(assignmentUUID)
	// 	},
	// )

	return binaryRelationFacade.Marshall(clone)
}

func (this *CommandGroupUser_Participated_AssignmentGroup_Relation) DispatchForLookupUnAssigned(
	tenantUUID uuid.UUID, assignmentUUID uuid.UUID,
) relation.IDBRelationInitiator {

	return this.DispatchDefault(tenantUUID, assignmentUUID)
}
