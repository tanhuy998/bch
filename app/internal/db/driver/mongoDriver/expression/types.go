package expression

type (
	IExpressionOperator interface {
		SetOperator(op string)
		GetOperator() string
	}

	IBinaryExpression[Left_T any, Right_T any] interface {
		IExpressionOperator
		SetLeftOperand(lhs Left_T)
		SetRightOperand(rhs Right_T)
		GetLeftOperand() Left_T
		GetRightOperand() Right_T
	}

	IUnaryExpression[Operand_T any] interface {
		IExpressionOperator
		SetOperand(operand Operand_T)
		GetOperand() Operand_T
	}
)
