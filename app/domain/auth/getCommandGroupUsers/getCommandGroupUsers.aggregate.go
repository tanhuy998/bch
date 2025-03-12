package getCommandGroupUsersDomain

import (
	"app/model"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/api/crud"
)

type (
	GetCommandGroupUserAggregate struct {
		aggregate.AggegateRoot[model.CommandGroupUser, model.CommandGroupUser]
		CommandGroupUser_User_Relation
	}
)

func (this *GetCommandGroupUserAggregate) MergeRelations() crud.IAggregateReader[model.CommandGroupUser] {

	// return this.AggregateRelations(
	// 	&this.CommandGroupUser_User_Relation,
	// )

	return this.ByDefaultRelations(
		&this.CommandGroupUser_User_Relation,
	)
}
