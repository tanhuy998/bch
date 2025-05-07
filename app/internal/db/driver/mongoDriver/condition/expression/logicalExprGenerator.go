package expression

// type (
// 	logical_expression_generator struct{}
// )

func NewLogicalExpressionGenerator() *logical_expression {

	return new(logical_expression)
}

// func (this *logical_expression_generator) Not() query.IConditionLogicalOperator {

// 	return &logical_expression{
// 		is_antonym: true,
// 	}
// }

// func (this *logical_expression_generator) And(
// 	conditionExpressions ...query.DataConditionMatchFunc,
// ) query.IDataConditionExpressionResult {

// 	return &logical_expression{
// 		op: "$and",
// 	}
// }

// func (this *logical_expression_generator) Or(
// 	conditionExpressions ...query.DataConditionMatchFunc,
// ) query.IDataConditionExpressionResult {

// 	return &logical_expression{
// 		op: "$or",
// 	}
// }
