package arithmetric

import "app/internal/db/driver/mongoDriver/filter/chain"

type (
	ArithmetricChain struct {
		chain.ExpressionChain
		initial_root_lhs interface{}
	}
)

func (this *ArithmetricChain) SetInitialOperand(val interface{}) {

	this.SetInitialRootRHS(val)
}

func (this *ArithmetricChain) SetInitialRootRHS(val interface{}) {

	if this.ExpressionChain.GetKeyNode() != nil {

		return
	}

	this.initial_root_lhs = val
}

func (this *ArithmetricChain) __getKeyNodeRHS() interface{} {

	switch keyNode := this.GetKeyNode(); {
	case keyNode == nil:
		return this.initial_root_lhs
	default:
		return keyNode.GetRightOperand()
	}
}

func (this *ArithmetricChain) Add(val interface{}) {

	this.Enchain(
		Add(
			this.__getKeyNodeRHS(), val,
		),
	)
}

func (this *ArithmetricChain) Substract(val interface{}) {

	this.Enchain(
		Substract(
			this.__getKeyNodeRHS(), val,
		),
	)
}

func (this *ArithmetricChain) Multiply(val interface{}) {

	this.Enchain(
		Multiply(
			this.__getKeyNodeRHS(), val,
		),
	)
}

func (this *ArithmetricChain) Divide(val interface{}) {

	this.Enchain(
		Divide(
			this.__getKeyNodeRHS(), val,
		),
	)
}

func (this *ArithmetricChain) Power(val interface{}) {

	this.Enchain(
		Power(
			this.__getKeyNodeRHS(), val,
		),
	)
}

func (this *ArithmetricChain) Logarit(base interface{}) {

	this.Enchain(
		Logarit(
			this.__getKeyNodeRHS(), base,
		),
	)
}
