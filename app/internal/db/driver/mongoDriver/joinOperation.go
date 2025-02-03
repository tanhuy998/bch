package mongoDriver

import (
	"app/internal/db/query"
)

type (
	join_op struct {
		From          string       `bson:"from"`
		Local_field   string       `bson:"localField"`
		Foreign_field string       `bson:"foreignField"`
		Alias         string       `bson:"as"`
		Pipeline      *mongo_query `bson:"pipeline,omitempty"`
		is_unwind     bool
	}
)

func (this *join_op) On(localField string, foreignField string) query.IJoinAlias {

	this.Local_field = localField
	this.Foreign_field = foreignField

	return this
}

func (this *join_op) As(name string) query.ISubQueryBuilder {

	this.Alias = name

	return this.Pipeline
}

func (this *join_op) NeedUnwind() bool {

	return this.is_unwind
}
