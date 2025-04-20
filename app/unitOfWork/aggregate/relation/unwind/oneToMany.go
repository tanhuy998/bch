package unwind

import (
	"app/internal/db/relation"
	aggregateRelation "app/unitOfWork/aggregate/relation"
)

type (
	OneToManyWith[Entity_T any] struct {
		aggregateRelation.OneToManyWith[Entity_T]
	}
)

func (this OneToManyWith[Entity_T]) ResolveRelation(
	local relation.IRelationLocalNavigator, foreignNavigator relation.IRelationForeignNavigator,
) {

	this.OneToManyWith.ResolveRelation(local, foreignNavigator)

	foreignNavigator.UnwindLocal()
}
