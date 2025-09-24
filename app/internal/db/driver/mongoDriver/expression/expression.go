package expression

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	Expression[Operand_T any] struct {
		expression_operator_t
		expression_operand_t[Operand_T]
	}
)

func (this *Expression[Operand_T]) AsBson() bson.D {

	return bson.D{
		{
			"$expr", this.BsonContent(),
		},
	}
}

func (this *Expression[Operand_T]) BsonContent() bson.D {

	return bson.D{
		{this.operator, this.expression_operand_t.operand},
	}
}

func (this Expression[Operand_T]) MarshalBSON() ([]byte, error) {

	return bson.Marshal(
		this.AsBson(),
	)
}

func (this Expression[Operand_T]) MarshalJSON() ([]byte, error) {

	return json.Marshal(
		this.AsBson(),
	)
}
