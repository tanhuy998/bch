package aggregate

import (
	"app/internal/db/query"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	assignmentGroupUser_assignmentGroup_Relation struct {
		aggregateRelation.OneToManyWith[model.AssignmentGroup]
	}
)

func (this assignmentGroupUser_assignmentGroup_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("assignmentGroupUUID", "uuid").As("assignmentGroups")
	}
}
