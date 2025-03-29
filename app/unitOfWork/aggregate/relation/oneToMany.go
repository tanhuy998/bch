package aggregateRelation

import (
	"app/internal/db/relation"
)

type (
	OneToManyWith[Foreign_Entity_T any] struct {
		abstract_relation[Foreign_Entity_T]
	}
)

func (this *OneToManyWith[Foreign_Entity_T]) ResolveRelation(
	refQueryBuilder relation.IRelationLocalNavigator, foreignInitializer relation.IRelationForeignNavigator,
) {

}

func (this *OneToManyWith[Foreign_Entity_T]) GetDBRelationKind() string {

	return "one_to_many"
}
