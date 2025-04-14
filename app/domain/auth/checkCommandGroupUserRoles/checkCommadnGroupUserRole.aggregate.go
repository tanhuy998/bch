package checkCommandGroupUserRoles

// import (
// 	"app/internal/db/relation"
// 	"app/model"
// 	"app/unitOfWork/aggregate"
// 	"app/unitOfWork/aggregate/api/crud"
// )

// type (
// 	CheckCommandGroupUserRoleAggregate struct {
// 		aggregate.DomainAggregateRoot[model.CommandGroupUserRole, model.CommandGroupUser]
// 		CommandGroupUser_CommandGroupUserRole_Relation
// 	}
// )

// func (this *CheckCommandGroupUserRoleAggregate) MergeRelation() crud.IAggregateReader[model.CommandGroupUserRole] {

// 	return this.ByRelations(
// 		this.CommandGroupUser_CommandGroupUserRole_Relation,
// 	)
// }

// func (this CheckCommandGroupUserRoleAggregate) ResolveRelation(
// 	local relation.IRelationLocalNavigator, foreignNavigator relation.IRelationForeignNavigator,
// ) {

// 	local.Unwind
// }
