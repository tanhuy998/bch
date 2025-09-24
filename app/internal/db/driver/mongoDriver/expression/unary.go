package expression

type (
	UnaryExpression[Operand_T any] struct {
		Expression[Operand_T]
		// expression_operator_t
		// operand Operand_T
	}
)

// func (this *UnaryExpression[Operand_T]) SetOperand(operand Operand_T) {

// 	this.operand = operand
// }

// func (this UnaryExpression[Operand_T]) MarshalBSON() ([]byte, error) {

// 	return bson.Marshal(
// 		bson.D{
// 			{this.operator, this.operand},
// 		},
// 	)
// }

// func (this UnaryExpression[Operand_T]) GetOperand() Operand_T {

// 	return this.operand
// }
