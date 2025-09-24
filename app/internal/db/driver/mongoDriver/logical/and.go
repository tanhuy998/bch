package logical

import (
	libCommon "app/internal/lib/common"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	AndExpression bson.A
)

func (this *AndExpression) AsBson() bson.D {

	return bson.D{
		{"$and", libCommon.PointerPrimitive(bson.A(*this))},
	}
}

func (this AndExpression) MarshalBSON() ([]byte, error) {

	switch len(this) {
	case 1:
		return bson.Marshal(this[0])
	default:
		return bson.Marshal(
			this.AsBson(),
		)
	}
}
