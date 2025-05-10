package innerJoin

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
	local relation.IRelationLocalNavigator, foreign relation.IRelationForeignNavigator,
) {

	this.OneToManyWith.ResolveRelation(local, foreign)
}

func (this OneToManyWith[Entity_T]) DetermineJoinOperation(
	initializer relation.IJoinDeterminerInitializer,
) relation.IJoinOperator {

	return initializer.AsInnerJoin()
}
