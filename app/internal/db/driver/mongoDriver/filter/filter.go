package filter

import (
	"app/internal/db/driver/mongoDriver/logical"
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	filter_t struct {
		logical.AndExpression
	}
)

func (this *filter_t) _manageField(fieldName string) *filter_expression_t {

	ret := NewFilterExpression(this, fieldName)

	this.AndExpression = append(this.AndExpression, ret)

	return ret
}

func (this *filter_t) Field(name string) query.INegatableDataConditionComparisonOperator /*query.IFilterExpressionOperator*/ {

	return this._manageField(name)
}

func (this filter_t) MarshalBSON() ([]byte, error) {

	return bson.Marshal(
		this.AndExpression.AsBson(),
	)
}

func (this filter_t) Get() bson.D {

	return this.AndExpression.AsBson()
}
