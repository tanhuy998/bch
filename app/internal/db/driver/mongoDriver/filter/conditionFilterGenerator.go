package filter

import (
	"app/internal/db/query"
)

type (
	ConditionFilterGenerator struct {
		filter_t //FilterGenerator
	}
)

func (this *ConditionFilterGenerator) Field(field string) query.INegatableDataConditionComparisonOperator {

	//ret := new(MongoNegationFilter)

	//ret.ref = &this.FilterGenerator
	//ret.lhs = field

	return this.filter_t._manageField(field)
}
