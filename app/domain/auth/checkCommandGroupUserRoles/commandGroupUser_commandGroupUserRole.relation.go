package checkCommandGroupUserRoles

import (
	"app/internal/db/query"
	"app/model"
	"app/unitOfWork/aggregate/relation/triat/leftJoin"
)

type (
	CommandGroupUser_CommandGroupUserRole_Relation struct {
		leftJoin.OneToManyWith[model.CommandGroupUserRole]
	}
)

func (this CommandGroupUser_CommandGroupUserRole_Relation) GetRelationInitFunc() query.JoinInitFunc {

	return func(join query.IJoinField) {

		join.On("uuid", "commandGroupUserUUID").As("roles")
	}
}
