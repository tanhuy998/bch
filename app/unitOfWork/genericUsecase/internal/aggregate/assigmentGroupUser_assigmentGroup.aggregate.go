package aggregate

import (
	"app/internal/db/query"
	"app/model"
	"app/unitOfWork/aggregate/relation/triat/leftJoin"
)

type (
	assignmentGroupUser_assignmentGroup_Relation struct {
		leftJoin.OneToManyWith[model.AssignmentGroup]
	}
)

func (this assignmentGroupUser_assignmentGroup_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("assignmentGroupUUID", "uuid").As("assignmentGroups")
	}
}
