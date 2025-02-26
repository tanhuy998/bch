package mongoDriver

import (
	"app/internal/db/driver/mongoDriver/lib"
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder"
	mongoRelation "app/internal/db/driver/mongoDriver/relation"
	"app/internal/db/query"
	"app/internal/db/relation"
	libCommon "app/internal/lib/common"
)

type (
	RelationDelegatorQueryBuilder[Entity_T any] struct {
		DelegatorQueryBuilder[Entity_T]
	}
)

func NewRelationDelegatorQueryBuilder[Entity_T any](det lib.IMongoDBCollection) *RelationDelegatorQueryBuilder[Entity_T] {

	ret := new(RelationDelegatorQueryBuilder[Entity_T])

	ret.delegator = det

	return ret
}

func (this *RelationDelegatorQueryBuilder[Entity_T]) _clone() *RelationDelegatorQueryBuilder[Entity_T] {

	var ret *RelationDelegatorQueryBuilder[Entity_T] = libCommon.PointerPrimitive(*this)

	ret.DelegatorQueryBuilder = *this.DelegatorQueryBuilder._clone()

	return ret
}

func (this *RelationDelegatorQueryBuilder[Entity_T]) Clone() query.IQueryBuilder[Entity_T] {

	return this._clone()
}

func (this *RelationDelegatorQueryBuilder[Entity_T]) AggregateRelations(
	relations ...relation.IDBRelationshipNavigator[mongoRelation.Query_Type],
) query.IQueryBuilder[Entity_T] {

	ret := this._clone()

	queryBuilder := (*mongoQueryBuilder.MongoRelationQueryBuilder)(&ret.MongoAggregateQueryBuilder)

	queryBuilder.PushRelations(relations...)

	return ret
}
