package filter

import (
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	filter_generator bson.D
)

func (this *filter_generator) reset() {

}

func (this *filter_generator) init() {

	if *this == nil {

		*this = filter_generator(bson.D{})
	}
}

func (this *filter_generator) Add(exprs ...bson.E) query.IFilterGenerator {

	this.init()

	*this = append(*this, exprs...)

	return this
}

func (this *filter_generator) Get() bson.D {

	return bson.D(*this)
}

func (this *filter_generator) Field(name string) query.IFilterExpressionOperator {

	return &MongoComparisonExprFilter{
		ref: this,
		lhs: name,
	}
}

func (this *filter_generator) GetCondtionExpression() interface{} {

	return (bson.D)(*this)
}
