package date

import (
	"app/internal/db/query"
)

type (
	date_iterator_t struct {
		date_operator_t
	}
)

func (this *date_iterator_t) Not() query.IFilterDateOperator {

	this.root.Negate()

	return this
}

func (this *date_iterator_t) Before() query.IFilterDateExtractionOperator {

	this.timeAspect = BEFORE

	return &this.date_operator_t
}

func (this *date_iterator_t) After() query.IFilterDateExtractionOperator {

	this.timeAspect = AFTER

	return &this.date_operator_t
}

func (this *date_iterator_t) Equal() query.IFilterDateExtractionOperator {

	this.timeAspect = EQUAL

	return &this.date_operator_t
}
