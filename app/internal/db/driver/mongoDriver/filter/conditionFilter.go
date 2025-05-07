package filter

import "go.mongodb.org/mongo-driver/bson"

type (
	MongoConditionExprFilter struct {
		MongoComparisonExprFilter
	}
)

func (this *MongoConditionExprFilter) EqualOneOfVals(vals ...interface{}) {

	this.In(vals...)
}

func (this *MongoConditionExprFilter) InNumericalIntegerRange(minVal int64, maxVal int64) {

	this._pushExpr(minVal, maxVal)
}

func (this *MongoConditionExprFilter) _pushExpr(minVal interface{}, maxVal interface{}) {

	expr := bson.E{
		this.lhs, bson.E{
			"$and", bson.A{
				bson.E{"$gte", minVal},
				bson.E{"$lte", maxVal},
			},
		},
	}

	switch this.is_antonym {
	case true:
		this.ref.Add(
			bson.E{
				"$not", expr,
			},
		)
	default:
		this.ref.Add(expr)
	}
}

func (this *MongoConditionExprFilter) InNumericalUnsignedIntergerRange(minVal uint64, maxVal uint64) {

	if minVal >= maxVal {

		panic("minVal must be less than max Val")
	}

	this._pushExpr(minVal, maxVal)
}

func (this *MongoConditionExprFilter) InNumericalFloatingPointRange(minVal float64, maxVal float64) {

	if minVal >= maxVal {

		panic("minVal must be less than max Val")
	}

	this._pushExpr(minVal, maxVal)
}
