package numeric

import (
	"app/internal/db/driver/mongoDriver/filter/operator"
	"app/internal/db/query"
)

type (
	numeric_operator_t struct {
		root *operator.SecondaryOperator
	}
)

func (this *numeric_operator_t) Not() query.IComparisonOperator {

	this.root.PrimaryOperator.Negate()

	return this
}

func (this *numeric_operator_t) EqualOneOfVals(vals ...interface{}) {

	this.root.SetOperator("$in")

	for i, val := range vals {

		switch v := val.(type) {
		case query.IQuerySelfReference:
			vals[i] = v.GetQuerySelfReference()
		}
	}

	this.root.SetRightOperand(vals)
}

func (this *numeric_operator_t) Equal(val interface{}) {

	this.root.SetOperator("$eq")
	this.root.SetRightOperand(val)
}

func (this *numeric_operator_t) GreaterThan(val interface{}) {

	this.root.SetOperator("$gt")
	this.root.SetRightOperand(val)
}

func (this *numeric_operator_t) GreaterOrEqual(val interface{}) {

	this.root.SetOperator("$gte")
	this.root.SetRightOperand(val)
}

func (this *numeric_operator_t) LessThan(val interface{}) {

	this.root.SetOperator("$lt")
	this.root.SetRightOperand(val)
}

func (this *numeric_operator_t) LessThanOrEqual(val interface{}) {

	this.root.SetOperator("$lte")
	this.root.SetRightOperand(val)
}

func (this *numeric_operator_t) In(vals ...interface{}) {

	this.root.SetOperator("$in")

	for i, val := range vals {

		switch v := val.(type) {
		case query.IQuerySelfReference:
			vals[i] = v.GetQuerySelfReference()
		}
	}

	this.root.SetRightOperand(vals)
}
