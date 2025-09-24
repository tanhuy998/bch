package condition

import (
	"app/internal/db/query"
	repositoryAPI "app/repository/api"
)

type (
	ConditionMaterializer struct {
		repositoryAPI.IPaginateQueryConditionUnit
	}
)

func (this *ConditionMaterializer) AsFilter(fn repositoryAPI.FilterFunc) query.IDataConditionExpressionResult {

	return &query_filter_operation_t{
		abstract_condition_t: abstract_condition_t{this.IPaginateQueryConditionUnit},
		fn:                   fn,
	}
}

func (this *ConditionMaterializer) AsConditionExpression(fn repositoryAPI.MatchFunc) query.IDataConditionExpressionResult {

	return &query_condition_expression_operation_t{
		abstract_condition_t: abstract_condition_t{this.IPaginateQueryConditionUnit},
		fn:                   fn,
	}
}
