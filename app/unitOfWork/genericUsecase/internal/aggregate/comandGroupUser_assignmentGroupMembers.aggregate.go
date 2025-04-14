package aggregate

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	CommandGroupUser_AssignmentGroupMember_Relation struct {
		aggregateRelation.OneToManyWith[model.AssignmentGroupMember]
		assignmentGroupUser_assignmentGroup_Relation
	}
)

func (this CommandGroupUser_AssignmentGroupMember_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {
		join.On("uuid", "commandGroupUserUUID").As("assignmentGroupMembers")
	}
}

func (this CommandGroupUser_AssignmentGroupMember_Relation) ResolveRelation(
	local relation.IRelationLocalNavigator, foreginNavigator relation.IRelationForeignNavigator,
) {

	this.OneToManyWith.ResolveRelation(local, foreginNavigator)

	foreginNavigator.Manipulate(
		func(foreign relation.IReadRelationQueryBuilder) {

			foreign.PushRelations(
				this.assignmentGroupUser_assignmentGroup_Relation,
			)
		},
	)
}
