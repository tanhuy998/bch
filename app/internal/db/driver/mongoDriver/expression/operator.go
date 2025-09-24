package expression

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	expression_operator_t struct {
		operator string
	}
)

func (this *expression_operator_t) SetOperator(op string) {

	this.operator = op
}

func (this *expression_operator_t) GetOperator() string {

	return this.operator
}

func (this expression_operator_t) MarshalBSON() ([]byte, error) {

	return bson.Marshal(this.operator)
}

func (this expression_operator_t) MarshalJSON() ([]byte, error) {

	return json.Marshal(this.operator)
}
