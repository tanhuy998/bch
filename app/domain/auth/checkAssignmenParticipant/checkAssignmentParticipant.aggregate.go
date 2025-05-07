package checkAssignmentParticipantDomain

import (
	"app/internal/db/query"
	"app/model"
	assignmentServicePort "app/port/assignment"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/api/crud"

	binaryRelation "app/unitOfWork/aggregate/relation/facade/binaryRelationFacade"
	binaryRelationTriat "app/unitOfWork/aggregate/relation/triat/binary"
)

type T struct {
	//AssignmentGroup model.AssignmentGroup `bson:"assignmentGroups"`
}

type (
	CheckAssignmentParticipantAggregate struct {
		aggregate.DomainAggregateRoot[T, model.CommandGroupUser]
		//pivot.Bridge[model.AssignmentGroupMember]
		//pivot.Foreign[model.AssignmentGroup]

		binaryRelationTriat.OneToManyBy[model.AssignmentGroupMember, model.AssignmentGroup]
	}
)

func (this CheckAssignmentParticipantAggregate) GetBridgeInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {
		join.On("uuid", "commandGroupUserUUID").As("assignmentGroupMembers")
	}
}

func (this CheckAssignmentParticipantAggregate) GetForeignInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {
		join.On("assignmentGroupUUID", "uuid").As("assignmentGroups")
	}
}

func (this *CheckAssignmentParticipantAggregate) MergeRelations(
	ctx assignmentServicePort.ICheckAssignmentParticipantContext,
) crud.IAggregateReader[T] {

	return this.MergeRelationsByDomain(
		ctx,
		binaryRelation.MarshallWithFilter(
			this,
			nil,
			func(foreignFilter query.IFilterExpression) {

				foreignFilter.Field("assignmentUUID").Equal(ctx.GetAssignmentUUID())
			},
		),
	)
}
