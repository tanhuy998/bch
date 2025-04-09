package getAssignmentGroupsDomain

import (
	"app/internal/db/query"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	AssigmentGroup_CommandGroup_Relation struct {
		aggregateRelation.OneToOneWith[model.CommandGroup]
	}
)

func (this AssigmentGroup_CommandGroup_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {
		join.On("commandGroupUUID", "uuid").As("commandGroup")
	}
}
