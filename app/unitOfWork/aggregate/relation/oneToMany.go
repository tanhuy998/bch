package aggregateRelation

import (
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/resolver"
)

type (
	OneToManyWith[Foreign_Entity_T any] struct {
		AbstractRelationInitiator[Foreign_Entity_T]
		resolver.One_To_Many
	}
)

func (this *OneToManyWith[Foreign_Entity_T]) ResolveRelation(
	refQueryBuilder relation.IRelationLocalNavigator, foreignInitializer relation.IRelationForeignNavigator,
) {

	this.ResolveRelation(
		refQueryBuilder, foreignInitializer,
	)
}

func (this *OneToManyWith[Foreign_Entity_T]) GetDBRelationKind() string {

	return "one_to_many"
}
