package mongoQueryBuilder

import (
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	JoinOperationInitializer struct {
		From                       string              `json:"from" bson:"from"`
		Local_field                string              `json:"localField" bson:"localField"`
		Foreign_field              string              `json:"foreignField" bson:"foreignField"`
		Alias                      string              `json:"as" bson:"as"`
		Pipeline                   *[]interface{}      `json:"pipeline,omitempty" bson:"pipeline,omitempty"`
		MongoAggregateQueryBuilder `json:"-" bson:"-"` // $lookup stage's pipeline, struct metadatas for bson are defined in mongo_pipeline struct
		is_unwind                  bool
	}
)

func (this *JoinOperationInitializer) initJoinPipeline() {

	if len(*this.Pipeline) > 0 {

		return
	}

	// if len(this.MongoAggregateQueryBuilder.P) == 0 {

	// 	this.MongoAggregateQueryBuilder.mongo_pipeline.init()
	// }

	this.Pipeline = &this.MongoAggregateQueryBuilder.P
}

func (this *JoinOperationInitializer) On(localField string, foreignField string) query.IJoinAlias {

	this.Local_field = localField
	this.Foreign_field = foreignField

	return this
}

func (this *JoinOperationInitializer) As(name string) query.IQueryBuilder {

	this.Alias = name

	return &this.MongoAggregateQueryBuilder
}

func (this *JoinOperationInitializer) NeedUnwind() bool {

	return this.is_unwind
}

func (this *JoinOperationInitializer) GetLocalField() string {

	return this.Local_field
}

func (this *JoinOperationInitializer) GetForeignField() string {

	return this.Foreign_field
}

func (this *JoinOperationInitializer) GetAliasName() string {

	return this.Alias
}

func (this *JoinOperationInitializer) SetLimit(num uint64) {

	// this.Pipeline.p = append(
	// 	[]interface{}{
	// 		bson.D{
	// 			{"$limit", num},
	// 		},
	// 	},
	// 	this.Pipeline.p...)

	this.MongoAggregateQueryBuilder.PrependStages(
		bson.D{
			{"$limit", num},
		},
	)
}

// func (this *JoinOperationInitializer) AggregateRelations(
// 	relations ...relation.IDBRelationInitiator[mongoRelation.Query_Type],
// ) query.IQueryBuilder {

// 	stage := MongoRelationQueryBuilder{}

// 	stage.PushRelations(relations...)

// 	this.MongoAggregateQueryBuilder.PushStages(
// 		stage,
// 	)

// 	return this
// }
