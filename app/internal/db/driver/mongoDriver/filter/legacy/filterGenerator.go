package filter

// import (
// 	"app/internal/db/query"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// type (
// 	FilterGenerator bson.D
// )

// func (this *FilterGenerator) reset() {

// }

// func (this *FilterGenerator) init() {

// 	if *this == nil {

// 		*this = FilterGenerator(bson.D{})
// 	}
// }

// func (this *FilterGenerator) Add(exprs ...bson.E) query.IFilterGenerator {

// 	this.init()

// 	*this = append(*this, exprs...)

// 	return this
// }

// func (this *FilterGenerator) Get() bson.D {

// 	return bson.D(*this)
// }

// func (this *FilterGenerator) Field(name string) query.IFilterExpressionOperator {

// 	return &MongoComparisonExprFilter{
// 		ref: this,
// 		lhs: name,
// 	}
// }

// func (this *FilterGenerator) GetCondtionExpression() interface{} {

// 	return (bson.D)(*this)
// }
