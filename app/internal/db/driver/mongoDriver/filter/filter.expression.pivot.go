package filter

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type (
	filter_expression_pivot_t struct {
		//filter_expression_evaluator_t
		parent    *filter_t //IExpressionList
		fieldName string
	}
)

func (this *filter_expression_pivot_t) AssertField(fieldName string, assertedType bsontype.Type) {
	/*
		prepend a field type assertion bson.D right before "this" expression
	*/

	this.parent.AndExpression = append(this.parent.AndExpression, this)

	lenFilter := len(this.parent.AndExpression)

	this.parent.AndExpression[lenFilter-2] = bson.D{
		{
			fieldName, bson.D{
				{"$type", assertedType},
			},
		},
	}
}

func (this filter_expression_pivot_t) CurrentFieldName() string {

	return this.fieldName
}

// func (f *filter_expression_pivot_t) SetOperator(operator string) {
// 	panic("TODO: Implement")
// }

func (f *filter_expression_pivot_t) OverrideOperand(operand interface{}) {
	panic("TODO: Implement")
}
