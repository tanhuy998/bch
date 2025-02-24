package mongoQueryBuilder

import (
	mongoRelation "app/internal/db/driver/mongoDriver/relation"
	"app/internal/db/relation"
)

type (
	/*
		Type Redefinition here just to add new method for the current struct,
		for better type convertion accross Relation delegator query builder and
		Delegator query builder.
	*/
	MongoRelationQueryBuilder MongoAggregateQueryBuilder
)

func (this *MongoRelationQueryBuilder) PushRelations(
	relations ...relation.IDBRelationshipNavigator[mongoRelation.Query_Type],
) {

	for _, rel := range relations {

		initializer := &JoinOperationInitializer{
			From: rel.GetDBStorageUnitName(),
		}
		initFunc := rel.GetRelationInitFunc()

		initFunc(initializer)

		this.PushStages(
			rel.ResolveRelationQuery(initializer)...,
		)
	}
}
