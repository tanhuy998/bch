package mongoQueryBuilder

import (
	mongoRelation "app/internal/db/driver/mongoDriver/relation"
	"app/internal/db/relation"
)

type (
	MongoRelationQueryBuilder struct {
		MongoAggregateQueryBuilder
	}
)

func (this *MongoRelationQueryBuilder) PushRelations(
	relations ...relation.IDBRelationshipQueryInitializer[mongoRelation.Query_Type],
) {

	for _, rel := range relations {

		initializer := &JoinOperationInitializer{
			From: rel.GetDBStorageUnitName(),
		}
		initFunc := rel.GetInitFunc()

		initFunc(initializer)

		this.MongoAggregateQueryBuilder.PushStages(
			rel.ResolveQuery(initializer)...,
		)
	}
}
