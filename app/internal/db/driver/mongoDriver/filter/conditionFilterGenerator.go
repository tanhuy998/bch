package filter

import (
	"app/internal/db/query"
)

type (
	condition_filter_generator struct {
		filter_generator
	}
)

func (this *condition_filter_generator) Field(field string) query.INegatableDataConditionComparisonOperator {

	ret := new(MongoNegationFilter)

	ret.ref = &this.filter_generator
	ret.lhs = field

	return ret
}
