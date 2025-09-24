package expression

import "go.mongodb.org/mongo-driver/bson"

type (
	binary_operand_pair_t[Left_Operand_T any, Right_Operand_T any] struct {
		lhs Left_Operand_T
		rhs Right_Operand_T
	}
)

func (this binary_operand_pair_t[Left_T, Right_T]) MarshalBSON() ([]byte, error) {

	return bson.Marshal(
		bson.A{this.lhs, this.rhs},
	)
}

type (
	BinaryExpression[Left_T any, Right_T any] struct {
		root Expression[binary_operand_pair_t[Left_T, Right_T]]
	}
)

func (this *BinaryExpression[Left_T, Right_T]) SetOperator(operator string) {

	this.root.SetOperator(operator)
}

func (this *BinaryExpression[Left_T, Right_T]) SetLeftOperand(lhs Left_T) {

	this.root.operand.lhs = lhs
}

func (this *BinaryExpression[Left_T, Right_T]) SetRightOperand(rhs Right_T) {

	this.root.operand.rhs = rhs
}

func (this BinaryExpression[Left_T, Right_T]) GetLeftOperand() Left_T {

	return this.root.operand.lhs
}

func (this BinaryExpression[Left_T, Right_T]) GetRightOperand() Right_T {

	return this.root.operand.rhs
}

func (this BinaryExpression[Left_T, Right_T]) MarshalBSON() ([]byte, error) {

	return this.root.MarshalBSON()
}
