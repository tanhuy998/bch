package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/relation"
)

type (
	MongoRelationQueryBuilder struct {
		queryBuilder.MongoAggregateQueryBuilder
		//query.IClonableQueryBuilder
		//relation_dispatcher
		relations []relation.IDBRelationInitiator
	}
)

func NewMongoRelationQueryBuilder() {

}

func (this *MongoRelationQueryBuilder) init() {

}

func (this *MongoRelationQueryBuilder) Init() {

	this.init()
}

func (this *MongoRelationQueryBuilder) _trackRelations(
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

	NewRelationDispatcher(
		&this.MongoAggregateQueryBuilder,
	)._dispatch(relations)

	this._trackRelations(relations)
}

func (this *MongoRelationQueryBuilder) Clone() relation.IClonableReadRelationQueryBuilder {

	return this._clone()
}

func (this *MongoRelationQueryBuilder) _clone() *MongoRelationQueryBuilder {

	this.init()

	ret := new(MongoRelationQueryBuilder)

	ret.MongoAggregateQueryBuilder = *this.MongoAggregateQueryBuilder.CloneThis()

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
