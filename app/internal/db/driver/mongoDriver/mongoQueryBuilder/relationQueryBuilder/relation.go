package relationQueryBuilder

import (
	"app/internal/db/driver/mongoDriver/mongoQueryBuilder/queryBuilder"
	"app/internal/db/relation"
)

type (
	MongoRelationQueryBuilder struct {
		queryBuilder.MongoAggregateQueryBuilder
		//query.IClonableQueryBuilder
		relations []relation.IDBRelationInitiator
	}
)

func NewMongoRelationQueryBuilder() {

}

func (this *MongoRelationQueryBuilder) init() {

	// if this.MongoAggregateQueryBuilder != nil {

	// 	return
	// }

	// this.MongoAggregateQueryBuilder = new(queryBuilder.MongoAggregateQueryBuilder)
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

// func (this *MongoRelationQueryBuilder) Transform(fn relation.RelationDataTransformFunc) {

// 	transformer := transform.NewRelationDataTransformer(
// 		this.relation_navigator,
// 	)

// 	fn(transformer)

// 	(*this.MongoAggregateQueryBuilder).PushStages(
// 		transformer.GetQuery()...,
// 	)
// }

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

	// for _, initiator := range relations {

	// 	initFn := initiator.GetRelationInitFunc()

	// 	if initFn == nil {

	// 		panic("relation init function could not be nil")
	// 	}

	// 	// initializer := &JoinOperationInitializer{
	// 	// 	From: initiator.GetDBStorageUnitName(),
	// 	// }

	// 	foreignNavigator := NewRelationForeignNavigator() // new(relation_navigator)
	// 	//initializer.stu_indentifier = initiator
	// 	foreignNavigator.join_op.From = initiator.GetDBStorageUnitName()

	// 	this.PushStages(
	// 		bson.D{
	// 			{"$lookup", &foreignNavigator.join_op},
	// 		},
	// 	)

	// 	initFn(&foreignNavigator.join_op)

	// 	initiator.ResolveRelation(
	// 		NewRelationLocalNaviagator(
	// 			//this.MongoAggregateQueryBuilder,
	// 			this, foreignNavigator,
	// 		),
	// 		foreignNavigator,
	// 	)

	// 	//foreignNavigator.join_op.Done()
	// }

	NewRelationDispatcher(
		&this.MongoAggregateQueryBuilder,
	).PushRelations(relations...)

	this._trackRelations(relations)
}

func (this *MongoRelationQueryBuilder) Clone() relation.IClonableReadRelationQueryBuilder {

	return this._clone()
}

func (this *MongoRelationQueryBuilder) _clone() *MongoRelationQueryBuilder {

	this.init()

	ret := new(MongoRelationQueryBuilder)

	// ret.MongoAggregateQueryBuilder = this.MongoAggregateQueryBuilder._clone()
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
