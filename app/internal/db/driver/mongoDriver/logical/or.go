package logical

import (
	libCommon "app/internal/lib/common"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	OrExpression bson.A
)

func (this *OrExpression) AsBson() bson.D {

	return bson.D{
		{"$or", libCommon.PointerPrimitive(bson.A(*this))},
	}
}

func (this OrExpression) MarshalBSON() ([]byte, error) {

	switch len(this) {
	case 1:
		return bson.Marshal(this[0])
	default:
		return bson.Marshal(
			this.AsBson(),
		)
	}
}
