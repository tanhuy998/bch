package resolver

import (
	"app/internal/db/relation"
)

type (
	One_To_One struct{}
)

func (this One_To_One) ResolveRelation(
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
