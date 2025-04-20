package resolver

import (
	"app/internal/db/relation"
)

type (
	Unwind struct {
	}
)

func (this Unwind) ResolveRelation(
	local relation.IRelationLocalNavigator, foreginNavigator relation.IRelationForeignNavigator,
) {

	foreginNavigator.UnwindLocal()
}
