package mongoRepositoryFilter

import (
	libCommon "app/internal/lib/common"
	repositoryAPI "app/repository/api"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	mongo_filter_expr struct {
		ref        MongoRepositoryFilterGenerator
		lhs        string
		rhs        interface{}
		is_antonym bool
	}
)

func (this *mongo_filter_expr) Equal(val interface{}) {

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

func (this *mongo_filter_expr) GreaterThan(val interface{}) {

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

func (this *mongo_filter_expr) GreaterOrEqual(val interface{}) {

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

func (this *mongo_filter_expr) LessThan(val interface{}) {

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

func (this *mongo_filter_expr) LessThanOrEqual(val interface{}) {

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

func (this *mongo_filter_expr) Or(fn repositoryAPI.FilterLogicalGroupFunc) {

	f :=

		fn()
}

func (f *mongo_filter_expr) And(fn repositoryAPI.FilterLogicalGroupFunc) {

	fn()
}

func (this *mongo_filter_expr) Not() repositoryAPI.IFilterExpressionOperator {

	this.is_antonym = true

	return this
}
