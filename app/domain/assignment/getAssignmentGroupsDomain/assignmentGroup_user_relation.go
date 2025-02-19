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

func (this *AssignmentGroup_User_Relation) GetInitFunc() query.JoinInitFunc {

	return func(queryBuilder query.IJoinField) {
		queryBuilder.On("createdBy", "uuid").As("createdUser")
	}
}
