package resolver

import (
	"app/internal/db/relation"
)

type (
	One_To_Many struct{}
)

func (this One_To_Many) ResolveRelation(
	local relation.IRelationLocalNavigator, foreign relation.IRelationForeignNavigator,
) {

}
