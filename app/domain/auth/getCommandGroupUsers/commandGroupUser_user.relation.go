package getCommandGroupUsersDomain

import (
	"app/internal/db/query"
	"app/model"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	CommandGroupUser_User_Relation struct {
		aggregateRelation.OneToOneWith[model.User]
	}
)

func (this *CommandGroupUser_User_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {
		join.On("userUUID", "uuid").As("user")
	}
}
