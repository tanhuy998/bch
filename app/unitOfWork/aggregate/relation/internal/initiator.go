package internal

import (
	"app/internal/db/query"
	"app/internal/db/relation"
	"app/unitOfWork/aggregate/relation/internal/api"
)

type (
	Initiator[
		Relation_Resolver_Kind_T api.IResolver,
		Initializer_Kind_T api.IInitializer,
	] struct {
		Resolver  Relation_Resolver_Kind_T
		Initiator Initializer_Kind_T
	}
)

func (this Initiator[R, I]) GetDBStorageUnitName() string {

	return this.Initiator.GetDBStorageUnitName()
}

func (this Initiator[R, I]) ResolveRelation(local relation.IRelationLocalNavigator, foreign relation.IRelationForeignNavigator) {

	this.Resolver.ResolveRelation(
		local, foreign,
	)
}

func (this Initiator[R, I]) GetDBRelationKind() string {

	return ""
}

func (this Initiator[R, I]) GetRelationInitFunc() query.JoinInitFunc {

	return this.Initiator.GetRelationInitFunc()
}
