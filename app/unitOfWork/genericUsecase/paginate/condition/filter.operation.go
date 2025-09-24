package condition

import repositoryAPI "app/repository/api"

type (
	query_filter_operation_t struct {
		abstract_condition_t
		fn repositoryAPI.FilterFunc
	}
)

func (this *query_filter_operation_t) ApplyConditionExpression() {

	this.AsFilter(this.fn)
}
