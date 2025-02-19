package mongoQueryBuilder

import (
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	JoinOperationInitializer struct {
		From          string                      `bson:"from"`
		Local_field   string                      `bson:"localField"`
		Foreign_field string                      `bson:"foreignField"`
		Alias         string                      `bson:"as"`
		Pipeline      *MongoAggregateQueryBuilder `bson:"pipeline,omitempty"`
		is_unwind     bool
	}
)

func (this *JoinOperationInitializer) On(localField string, foreignField string) query.IJoinAlias {

	this.Local_field = localField
	this.Foreign_field = foreignField

	return this
}

func (this *JoinOperationInitializer) As(name string) query.ISubQueryBuilder {

	this.Alias = name

	return this.Pipeline
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

	this.Pipeline.p = append(
		[]interface{}{
			bson.D{
				{"$limit", num},
			},
		},
		this.Pipeline.p...)
}
