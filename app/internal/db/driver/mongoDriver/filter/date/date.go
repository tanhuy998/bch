package date

import (
	"app/internal/db/driver/mongoDriver/filter/operator"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	filter_expression_date_casting_t struct {
		date_iterator_t
	}
)

func NewFilterExpressionDateCasting(
	operator *operator.SecondaryOperator,
) *filter_expression_date_casting_t {

	ret := new(filter_expression_date_casting_t)

	ret.root = operator
	ret.root.SetAssertedType(bson.TypeDateTime)

	return ret
}
