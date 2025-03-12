package getSingleAssignmentDomain

import (
	"app/model"
	"app/unitOfWork/aggregate"
	"app/unitOfWork/aggregate/api/crud"
)

type (
	GetSingleAssigmentAggregate struct {
		aggregate.AggegateRoot[model.Assignment, model.Assignment]
		Assignment_User_Relation
	}
)

func (this *GetSingleAssigmentAggregate) MergeRelations() crud.IAggregateReader[model.Assignment] {

	return this.ByDefaultRelations(
		&this.Assignment_User_Relation,
	)
}
