package aggregateRelation

import (
	"app/internal/db/relation"
)

type (
	Query_Type = []interface{}
)

type (
	OneToOneWith[Foreign_Entity_T any] struct {
		abstract_relation[Foreign_Entity_T]
	}
)

func (this *OneToOneWith[Foreign_Entity_T]) ResolveRelation(
	local relation.IRelationLocalNavigator, foreign relation.IRelationForeignNavigator,
) {

	foreign.Manipulate(
		func(foreign relation.IReadRelationQueryBuilder) {

			foreign.Limit(1)
		},
	)

	local.Transform(
		func(transform relation.IRelationLocalDataTransformer) {

			alias := foreign.GetAliasName()

			transform.Set(alias).AsForeign().FirstElement()
		},
	)
}

/*
for debug log
*/
func (this *OneToOneWith[Foreign_Entity_T]) GetDBRelationKind() string {

	return "one_to_one"
}
