package comparision

import (
	"app/internal/db/query"
)

type (
	ComparisionExpression struct {
		abstract_comparision_t
		//is_antonym bool
	}
)

// func NewComparisonExpression(fieldName string) *ComparisionExpression {

// 	ret := new(ComparisionExpression)

// 	ret.SetCurrentFieldName(fieldName)

// 	ret.Init()

// 	return ret
// }

func (this *ComparisionExpression) Equal(val interface{}) {

	this.prepare()

	this.SetOperator("$eq")
	this.SetRightOperand(val)
}

func (this *ComparisionExpression) GreaterThan(val interface{}) {

	this.prepare()

	this.SetOperator("$gt")
	this.SetRightOperand(val)
}

func (this *ComparisionExpression) GreaterOrEqual(val interface{}) {

	this.prepare()

	this.SetOperator("$gte")
	this.SetRightOperand(val)
}

func (this *ComparisionExpression) LessThan(val interface{}) {

	this.prepare()

	this.SetOperator("$lt")
	this.SetRightOperand(val)
}

func (this *ComparisionExpression) LessThanOrEqual(val interface{}) {

	this.prepare()

	this.SetOperator("$lte")
	this.SetRightOperand(val)
}

func (this *ComparisionExpression) In(vals ...interface{}) {

	this.prepare()

	this.SetOperator("$in")
	this.SetRightOperand(vals)
}

func (this *ComparisionExpression) Not() query.IDataConditionComparisonOperator /*query.IComparisonOperator*/ {

	this.Negate()

	return this
}

func (this *ComparisionExpression) EqualOneOfVals(vals ...interface{}) {

	this.In(vals...)
}

// func (this *ComparisionExpression) Negate() {

// 	this.is_antonym = true
// }

// func (this ComparisionExpression) MarshalBSON() ([]byte, error) {

// 	switch this.is_antonym {
// 	case true:
// 		return logical.NewNotLogicalExpression(this.Expression).MarshalBSON()
// 	default:
// 		return this.Expression.MarshalBSON()
// 	}
// }
