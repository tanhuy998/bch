package mongoQueryBuilder

import (
	"app/internal/db/relation"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	/*
		Type Redefinition here just to add new method for the current struct,
		for better type convertion accross Relation delegator query builder and
		Delegator query builder.
	*/
	MongoRelationQueryBuilder struct {
		MongoAggregateQueryBuilder
		relations []relation.IDBRelationInitiator
	}
)

func (this *MongoRelationQueryBuilder) _pushRelations(
	relations []relation.IDBRelationInitiator,
) {

	if len(this.relations) == 0 {

		this.relations = relations
		return
	}

	this.relations = append(this.relations, relations...)
}

func (this *MongoRelationQueryBuilder) PushRelations(
	relations ...relation.IDBRelationInitiator,
) {

	// for _, rel := range relations {

	// 	initializer := &JoinOperationInitializer{
	// 		From: rel.GetDBStorageUnitName(),
	// 	}
	// 	initFunc := rel.GetRelationInitFunc()

	// 	initFunc(initializer)

	// 	this.PushStages(
	// 		rel.ResolveRelationQuery(initializer)...,
	// 	)
	// }

	for _, initiator := range relations {

		initFn := initiator.GetRelationInitFunc()

		if initFn == nil {

			panic("relation init function could not be nil")
		}

		initializer := &JoinOperationInitializer{
			From: initiator.GetDBStorageUnitName(),
		}

		initFn(initializer)

		this.PushStages(
			bson.D{
				{"$lookup", initializer},
			},
		)

		initiator.ResolveRelation(
			&this.MongoAggregateQueryBuilder, initializer,
		)
	}

	this._pushRelations(relations)
}

func (this *MongoRelationQueryBuilder) Clone() relation.IClonableReadRelationQueryBuilder {

	return this._clone()
}

func (this *MongoRelationQueryBuilder) _clone() *MongoRelationQueryBuilder {

	ret := new(MongoRelationQueryBuilder)

	ret.MongoAggregateQueryBuilder = *this.MongoAggregateQueryBuilder._clone()

	ret.relations = make([]relation.IDBRelationInitiator, len(this.relations))
	copy(ret.relations, this.relations)

	return ret
}

func (this *MongoRelationQueryBuilder) GetDetailQueryDebugLog() interface{} {

	ret := aggregate_relation_debug_log{
		Relations: make([]relation_debug_log, len(this.relations)),
	}

	for i, initiator := range this.relations {

		ret.Relations[i] = relation_debug_log{
			RelationType:      initiator.GetDBRelationKind(),
			ForeignCollection: initiator.GetDBStorageUnitName(),
		}
	}

	return ret
}
