package getAssignmentGroupsDomain

import (
	"app/internal/db/driver/mongoDriver/mongoStorage"
	"app/internal/db/query"
	"app/model"
	"app/repository"
	"app/unitOfWork/aggregate"
	"fmt"
)

type (
	GetAssignmentGroupAggegate struct {
		aggregate.AggegateRoot[model.AssignmentGroup, model.AssignmentGroup]
		AssignmentGroup_User_Relation
		AssigmentGroup_CommandGroup_Relation
		Stu              mongoStorage.IMongoDBStorageUnit[model.User]
		UserRepo         repository.IUser
		CommandGroupRepo repository.ICommandGroup
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

	fmt.Println("test storage unit:", this.Stu.GetDBStorageUnitName())

	return this.AggregateRelationsOnceOrDefault(
		&this.AssignmentGroup_User_Relation,
		&this.AssigmentGroup_CommandGroup_Relation,
	)
}
