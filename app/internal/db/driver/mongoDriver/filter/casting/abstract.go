package casting

// import (
// 	"app/internal/db/driver/mongoDriver/filter/casting/iterator"
// 	"app/internal/db/query"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/bson/bsontype"
// )

// type (
// 	abstract_casting_expression_t struct {
// 		iterator.IIterator
// 		parent             IFilterExpressionPivot //IFilterExpression
// 		asserted_bson_type bsontype.Type
// 	}
// )

// func (this *abstract_casting_expression_t) init() {

// 	this.parent.AssertField(
// 		this.parent.CurrentFieldName(), this.asserted_bson_type,
// 	)
// }

// func (this abstract_casting_expression_t) AssertFieldIfSelfReference(
// 	fieldName interface{},
// ) (referenceFieldName string) {

// 	switch v := fieldName.(type) {
// 	case query.IQuerySelfReference:
// 		this.parent.AssertField(v.GetQuerySelfReference(), this.asserted_bson_type)
// 		return fmt.Sprintf("$%s", v.GetQuerySelfReference())
// 	default:
// 		return ""
// 	}
// }
