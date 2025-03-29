package queryBuilder

import (
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	JoinOperationInitializer struct {
		From          string       `json:"from" bson:"from"`
		Local_field   string       `json:"localField" bson:"localField"`
		Foreign_field string       `json:"foreignField" bson:"foreignField"`
		Alias         string       `json:"as" bson:"as"`
		Pipeline      JoinPipeline `json:",inline" bson:",inline"`
		is_unwind     bool
		raw           *bson.M
	}
)

func NewJoinOperationQueryBuilder() *JoinOperationInitializer {

	ret := new(JoinOperationInitializer)

	//ret.Init()

	ret.Pipeline.Init()

	return ret
}

func (this *JoinOperationInitializer) On(localField string, foreignField string) query.IJoinAlias {

	this.Local_field = localField
	this.Foreign_field = foreignField

	return this
}

func (this *JoinOperationInitializer) As(name string) query.IQueryBuilder {

	this.Alias = name

	// return &this.MongoAggregateQueryBuilder

	return &(&this.Pipeline).MongoAggregateQueryBuilder
}

func (this *JoinOperationInitializer) NeedUnwind() bool {

	return this.is_unwind
}

func (this *JoinOperationInitializer) Done() {

	this.Pipeline.Done()

	switch this.raw {
	case nil:
		this.raw = &bson.M{
			"from":         this.From,
			"localField":   this.Local_field,
			"foreignField": this.Foreign_field,
			"as":           this.Alias,
			"pipeline":     this.Pipeline.MongoAggregateQueryBuilder.P,
		}
	default:
		(*this.raw)["pipeline"] = this.Pipeline.MongoAggregateQueryBuilder.P
		(*this.raw)["as"] = this.Alias
	}
}

func (this *JoinOperationInitializer) GetRawQuery() *bson.M {

	this.Done()

	return this.raw
}

// func (this *JoinOperationInitializer) GetLocalField() string {

// 	return this.Local_field
// }

// func (this *JoinOperationInitializer) GetForeignField() string {

// 	return this.Foreign_field
// }

// func (this *JoinOperationInitializer) GetAliasName() string {

// 	return this.Alias
// }

// func (this *JoinOperationInitializer) SetLimit(num uint64) {

// 	// this.Pipeline.p = append(
// 	// 	[]interface{}{
// 	// 		bson.D{
// 	// 			{"$limit", num},
// 	// 		},
// 	// 	},
// 	// 	this.Pipeline.p...)

// 	this.MongoAggregateQueryBuilder.PrependStages(
// 		bson.D{
// 			{"$limit", num},
// 		},
// 	)
// }

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
