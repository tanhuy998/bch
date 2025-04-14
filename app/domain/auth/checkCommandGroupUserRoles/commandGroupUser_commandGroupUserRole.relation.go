package checkCommandGroupUserRoles

// import (
// 	"app/internal/db/query"
// 	"app/model"
// 	aggregateRelation "app/unitOfWork/aggregate/relation"
// )

// type (
// 	CommandGroupUser_CommandGroupUserRole_Relation struct {
// 		aggregateRelation.OneToManyWith[model.CommandGroupUserRole]
// 	}
// )

// func (this CommandGroupUser_CommandGroupUserRole_Relation) GetRelationInitFunc() query.JoinInitFunc {

// 	return func(join query.IJoinField) {

// 		join.On("uuid", "commandGroupUserUUID").As("roles")
// 	}
// }
