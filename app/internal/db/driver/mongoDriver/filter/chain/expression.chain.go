package chain

type (
	ExpressionChain struct {
		root *ExpressionChainNode
	}
)

func (this *ExpressionChain) GetKeyNode() *ExpressionChainNode {

	return this.root
}

func (this *ExpressionChain) Enchain(newExpression *ExpressionChainNode) {

	if newExpression == nil {

		return
	}

	switch this.root {
	case nil:
		this.root = newExpression
	default:
		newExpression.SetLeftOperand(this.root)
		this.root = newExpression
	}
}

func (this ExpressionChain) MarshalBSON() ([]byte, error) {

	return this.root.MarshalBSON()
}
