package filter

import (
	"app/internal/db/query"
	libCommon "app/internal/lib/common"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	MongoComparisonExprFilter struct {
		ref        *FilterGenerator
		lhs        string
		rhs        interface{}
		is_antonym bool
	}
)

func (this *MongoComparisonExprFilter) Equal(val interface{}) {

	if this.lhs == "" {

		return
	}

	if this.is_antonym {

		this.ref.Add(
			bson.E{
				this.lhs, bson.E{
					"$ne", val,
				},
			},
		)
		return
	}

	this.ref.Add(bson.E{this.lhs, val})
}

func (this *MongoComparisonExprFilter) GreaterThan(val interface{}) {

	if this.lhs == "" {

		return
	}

	op := libCommon.Ternary(this.is_antonym, "$lte", "$gt")

	this.ref.Add(
		bson.E{
			this.lhs, bson.E{
				op, val,
			},
		},
	)
}

func (this *MongoComparisonExprFilter) GreaterOrEqual(val interface{}) {

	if this.lhs == "" {

		return
	}

	op := libCommon.Ternary(this.is_antonym, "$lt", "$gte")

	this.ref.Add(
		bson.E{
			this.lhs, bson.E{
				op, val,
			},
		},
	)
}

func (this *MongoComparisonExprFilter) LessThan(val interface{}) {

	if this.lhs == "" {

		return
	}

	op := libCommon.Ternary(this.is_antonym, "$gte", "$lt")

	this.ref.Add(
		bson.E{
			this.lhs, bson.E{
				op, val,
			},
		},
	)
}

func (this *MongoComparisonExprFilter) LessThanOrEqual(val interface{}) {

	if this.lhs == "" {

		return
	}

	op := libCommon.Ternary(this.is_antonym, "$gt", "$lte")

	this.ref.Add(
		bson.E{
			this.lhs, bson.E{
				op, val,
			},
		},
	)
}

func (this *MongoComparisonExprFilter) In(vals ...interface{}) {

	op := libCommon.Ternary(this.is_antonym, "$nin", "$in")

	this.ref.Add(
		bson.E{
			this.lhs, bson.E{
				op, vals,
			},
		},
	)
}

func (this *MongoComparisonExprFilter) Not() query.IComaparisonOperator {

	this.is_antonym = true

	return this
}
