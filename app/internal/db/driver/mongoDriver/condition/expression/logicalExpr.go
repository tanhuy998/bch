package expression

import (
	"app/internal/db/query"

	"go.mongodb.org/mongo-driver/bson"
)

type (
	logical_expression struct {
		ConditionExpressionInitializer
		refConditionExpression *bson.D
		elements               []bson.D
		op                     string
		is_antonym             bool
	}
)

func (this *logical_expression) _resolveActualOpIfNegation() string {

	if !this.is_antonym {

		return this.op
	}

	return "$nor"
}

func (this *logical_expression) GetCondtionExpression() interface{} {

	// ret := bson.D{
	// 	{this.op, this.elements},
	// }

	// switch this.is_antonym {
	// case true:
	// 	return bson.D{
	// 		{"$not", ret},
	// 	}
	// default:
	// 	return ret
	// }

	return bson.D{
		{this._resolveActualOpIfNegation(), this.elements},
	}
}

func (this *logical_expression) ApplyConditionExpression() {

	switch len(this.elements) {
	case 0:
		return
	case 1:
		*this.refConditionExpression = this.elements[0]
	default:
		*this.refConditionExpression = bson.D{
			{this._resolveActualOpIfNegation(), this.elements},
		}
	}
}

func (this *logical_expression) _resolve(
	expresstionFuncs []query.DataConditionMatchFunc,
) {

	this.elements = make([]bson.D, 0)

	for _, fn := range expresstionFuncs {

		if fn == nil {

			panic("logical expression function must not be nil")
		}

		// res := fn(&this.ConditionExpressionInitializer)

		// this.elements[i] = res.GetCondtionExpression()

		var expression bson.D

		conditionExpressInitializer := NewConditionExpressionInitializer(&expression)

		result := fn(conditionExpressInitializer)

		if result == nil {

			continue
		}

		switch result.ApplyConditionExpression(); {
		case expression == nil:
			continue
		default:
			this.elements = append(this.elements, expression)
		}
	}
}

func (this *logical_expression) Not() query.IConditionLogicalOperator {

	this.is_antonym = true

	return this
}

func (this *logical_expression) And(conditionExpressions ...query.DataConditionMatchFunc) query.IDataConditionExpressionResult {

	this.op = "$and"

	this._resolve(conditionExpressions)

	return this
}

func (this *logical_expression) Or(conditionExpressions ...query.DataConditionMatchFunc) query.IDataConditionExpressionResult {

	this.op = "$or"

	//this.op = libCommon.Ternary(!this.is_antonym, "$or", "$nor")

	this._resolve(conditionExpressions)

	return this
}
