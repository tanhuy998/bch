package getSingleAssignmentDomain

import (
	"app/internal/db/query"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	Assignment_User_Relation struct {
		aggregateRelation.OneToOneWith[model.User]
	}
)

// Implement "app/unitOfWork/aggregate/relation".IDBRelationshipDeclarativeInitializer interface
func (this *Assignment_User_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("createdBy", "uuid").As("createdUser")
	}
}
