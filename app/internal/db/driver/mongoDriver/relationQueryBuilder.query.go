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
		//DelegatorQueryBuilder[Entity_T]
		//MongoDBDelegator[Entity_T, Source_Repo_Entity]
		DelegatorQueryBuilder[Entity_T]
	}
)

func NewRelationDelegatorQueryBuilder[Entity_T any](det lib.IMongoDBCollection) *RelationDelegatorQueryBuilder[Entity_T] {

	ret := new(RelationDelegatorQueryBuilder[Entity_T])

	ret.delegator = det

	return ret
}

func (this *RelationDelegatorQueryBuilder[Entity_T]) Clone() query.IQueryBuilder[Entity_T] {

	return libCommon.PointerPrimitive(*this)
}

func (this *RelationDelegatorQueryBuilder[Entity_T]) AggregateRelations(
	relations ...relation.IDBRelationshipNavigator[mongoRelation.Query_Type],
) query.IQueryBuilder[Entity_T] {

	queryBuilder := mongoQueryBuilder.MongoRelationQueryBuilder(this.MongoAggregateQueryBuilder)

	queryBuilder.PushRelations(relations...)

	return this.Clone()
}
