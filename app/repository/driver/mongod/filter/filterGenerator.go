package mongoRepositoryFilter

// import (
// 	"app/internal/db/query"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// type (
// 	MongoRepositoryFilterGenerator bson.D
// )

// func (this *MongoRepositoryFilterGenerator) reset() {

// }

// func (this *MongoRepositoryFilterGenerator) init() {

// 	if *this == nil {

// 		*this = MongoRepositoryFilterGenerator(bson.D{})
// 	}
// }

// func (this *MongoRepositoryFilterGenerator) Add(exprs ...bson.E) query.IFilterGenerator /* repositoryAPI.IFilterGenerator */ {

// 	this.init()

// 	*this = append(*this, exprs...)

// 	return this
// }

// func (this *MongoRepositoryFilterGenerator) Get() bson.D {

// 	switch *this {
// 	case nil:
// 		return bson.D{}
// 	default:
// 		return bson.D(*this)
// 	}
// }

// func (this *MongoRepositoryFilterGenerator) Field(name string) query.IFilterExpressionOperator /* repositoryAPI.IFilterExpressionOperator */ {

// 	return &mongo_filter_expr{
// 		ref: this,
// 		lhs: name,
// 	}
// }
