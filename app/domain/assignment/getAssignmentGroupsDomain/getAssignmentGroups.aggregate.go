package getAssignmentGroupsDomain

import (
	"app/internal/db/query"
	"app/model"
	"app/unitOfWork/aggregate"
)

type (
	GetAssignmentGroupAggegate struct {
		aggregate.AggegateRoot[model.AssignmentGroup, model.AssignmentGroup]
		AssignmentGroup_User_Relation
		AssigmentGroup_CommandGroup_Relation
	}
)

func (this *GetAssignmentGroupAggegate) GetAssignentGroups() query.IQueryBuilder[model.AssignmentGroup] {

	// return this.Query().
	// 	Join(this.UserRepo, func(queryBuilder query.IJoinField) {
	// 		queryBuilder.On("createdBy", "uuid").As("createdUser").Limit(1)
	// 	}).
	// 	Join(this.CommandGroupRepo, func(queryBuilder query.IJoinField) {
	// 		queryBuilder.On("commandGroupUUID", "uuid").As("commandGroup").Limit(1)
	// 	}).
	// 	ExcludeFields(
	// 		"user.password",
	// 		"user.secret",
	// 	)

	return this.AggregateRelationsOnceOrDefault(
		&this.AssignmentGroup_User_Relation,
		&this.AssigmentGroup_CommandGroup_Relation,
	)
}
