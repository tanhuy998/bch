package getSingleAssignmentDomain

import (
	"app/internal/db/query"
	"app/model"
	"app/unitOfWork/aggregate"
)

type (
	GetSingleAssigmentAggregate struct {
		aggregate.AggegateRoot[model.Assignment, model.Assignment]
		Assignment_User_Relation
	}
)

func (this *GetSingleAssigmentAggregate) MergeRelations() query.IQueryBuilder[model.Assignment] {

	return this.AggregateRelations(
		&this.Assignment_User_Relation,
	)
}
