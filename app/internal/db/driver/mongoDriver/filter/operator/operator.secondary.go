package operator

import (
	"app/internal/db/driver/mongoDriver/expression"

	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type (
	SecondaryOperator struct {
		PrimaryOperator
		root               IFilterExpressionPivot
		asserted_bson_type bsontype.Type
	}
)

func (this *SecondaryOperator) Init() {

	this.root.AssertField(
		this.root.CurrentFieldName(), this.asserted_bson_type,
	)
}

func (this SecondaryOperator) Pivot() IFilterExpressionPivot {

	return this.root
}

func (this SecondaryOperator) GetBasedExpression() expression.IBinaryExpression[interface{}, interface{}] {

	return &this.base_t
}

func (this *SecondaryOperator) SetAssertedType(t bsontype.Type) {

	this.asserted_bson_type = t
}

func (this *SecondaryOperator) SetRoot(p IFilterExpressionPivot) {

	this.root = p
}

func (this *SecondaryOperator) OverrideExpression() IExpressionEvaluator {

	return &this.expression_evaluator_t
}

func (this *SecondaryOperator) Root() IFilterExpression {

	return this.root
}

func (this SecondaryOperator) AssertFieldIfSelfReference(
	fieldName interface{},
) (referenceFieldName string) {

	switch v := fieldName.(type) {
	case query.IQuerySelfReference:
		this.root.AssertField(v.Origin(), this.asserted_bson_type)
		return v.GetQuerySelfReference()
	default:
		return ""
	}
}
