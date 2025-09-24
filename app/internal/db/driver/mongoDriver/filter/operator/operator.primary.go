package operator

import (
	"go.mongodb.org/mongo-driver/bson"
)

type (
	PrimaryOperator struct {
		expression_evaluator_t
		is_antonym bool
	}
)

func (this PrimaryOperator) MarshalBSON() ([]byte, error) {

	switch this.is_antonym {
	case true:
		return bson.Marshal(
			bson.D{
				{
					"$expr", bson.D{
						{"$not", this.base_t.BsonContent()},
					},
				},
			},
		)
	default:
		return this.base_t.MarshalBSON()
	}
}

func (this *PrimaryOperator) Negate() {

	this.is_antonym = true
}
