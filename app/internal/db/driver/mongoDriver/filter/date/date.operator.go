package date

import (
	"app/internal/db/driver/mongoDriver/filter/operator"
	"fmt"
	"time"
)

type (
	time_comparison_aspect_t int
)

const (
	EQUAL time_comparison_aspect_t = iota
	AFTER
)

const (
	BEFORE time_comparison_aspect_t = EQUAL - 1
)

type (
	date_operator_t struct {
		root       *operator.SecondaryOperator
		timeAspect time_comparison_aspect_t
	}
)

func (this *date_operator_t) getTimeAspect() string {
	switch this.timeAspect {
	case AFTER:
		return "$gt"
	case BEFORE:
		return "$lt"
	default:
		return "$eq"
	}
}

func (this *date_operator_t) prepareBinaryExpresionQuery() *time.Time {

	expr := this.root.OverrideExpression().AsBinaryExpression()
	expr.SetLeftOperand(
		fmt.Sprintf(`$%s`, this.root.Pivot().CurrentFieldName()),
	)
	expr.SetOperator(
		this.getTimeAspect(),
	)
	ret := &time.Time{}
	expr.SetRightOperand(ret)

	return ret
}

func (this *date_operator_t) prepareSelfReferenceQuery(refField string) *self_ref_time_comparator_t {

	expr := this.root.OverrideExpression().AsUnaryExpression()
	expr.SetOperator("$and")

	operand := &self_ref_time_comparator_t{
		current_field: this.root.Pivot().CurrentFieldName(),
		ref_field:     refField,
		aspect:        this.getTimeAspect(),
	}

	expr.SetOperand(operand)

	return operand
}

func (this *date_operator_t) Date(date interface{}) {

	switch self := this.root.AssertFieldIfSelfReference(date); {
	case self != "":
		this.root.SetRightOperand(self)
		return
	}

	switch t := date.(type) {
	case time.Time, *time.Time:
		this.root.SetRightOperand(t)
	default:
		panic("wrong type of date in date query filter")
	}
}

func (this *date_operator_t) YearOf(date interface{}) {

	switch self := this.root.AssertFieldIfSelfReference(date); {
	case self != "":
		comparator := this.prepareSelfReferenceQuery(self)
		comparator.includeYear()
		return
	}

	var input time.Time

	switch t := date.(type) {
	case time.Time:
		input = t
	case *time.Time:
		input = *t
	default:
		panic("wrong type of date in date query filter")
	}

	queriedTime := this.prepareBinaryExpresionQuery()
	*queriedTime = time.Date(
		input.Year(), 0, 0, 0, 0, 0, 0, input.Location(),
	)
}

func (this *date_operator_t) MonthOf(date interface{}) {

	switch self := this.root.AssertFieldIfSelfReference(date); {
	case self != "":
		this.prepareSelfReferenceQuery(self).includeMonth()
		return
	}

	var input time.Time

	switch t := date.(type) {
	case time.Time:
		input = t
	case *time.Time:
		input = *t
	default:
		panic("wrong type of date in date query filter")
	}

	queriedTime := this.prepareBinaryExpresionQuery()
	*queriedTime = time.Date(
		input.Year(), input.Month(), 0, 0, 0, 0, 0, input.Location(),
	)
}

func (this *date_operator_t) DateOf(date interface{}) {

	switch self := this.root.AssertFieldIfSelfReference(date); {
	case self != "":
		this.prepareSelfReferenceQuery(self).includeDay()
		return
	}

	var input time.Time

	switch t := date.(type) {
	case time.Time:
		input = t
	case *time.Time:
		input = *t
	default:
		panic(
			fmt.Sprintf("wrong type of date in date query, %T given", t),
		)
	}

	queriedTime := this.prepareBinaryExpresionQuery()
	*queriedTime = time.Date(
		input.Year(), input.Month(), input.Day(), 0, 0, 0, 0, input.Location(),
	)
}
