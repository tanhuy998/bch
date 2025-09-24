package condition

import repositoryAPI "app/repository/api"

type (
	query_condition_expression_operation_t struct {
		abstract_condition_t
		fn repositoryAPI.MatchFunc
	}
)

func (this *query_condition_expression_operation_t) ApplyConditionExpression() {

	this.AsConditionExpression(this.fn)
}
