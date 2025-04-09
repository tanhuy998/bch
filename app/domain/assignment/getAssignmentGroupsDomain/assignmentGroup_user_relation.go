package getAssignmentGroupsDomain

import (
	"app/internal/db/query"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	AssignmentGroup_User_Relation struct {
		aggregateRelation.OneToOneWith[model.User]
	}
)

func (this AssignmentGroup_User_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {
		join.On("createdBy", "uuid").As("createdUser")
	}
}
