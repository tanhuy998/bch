package iterator

import "app/internal/db/driver/mongoDriver/filter/operator"

type (
	IEnchainable_T[Node_T any] interface {
		SetInitialOperand(val interface{})
		Enchain(*Node_T)
		GetKeyNode() *Node_T
	}
)

type (
	Iterator[
		Chain_T IEnchainable_T[Chain_Node_T],
		Chain_Node_T any,
	] struct {
		root                    *operator.SecondaryOperator
		root_iterator_delegator ExpressionDelegator[Chain_T, interface{}]
	}
)

func (this *Iterator[Chain_T, Node_T]) Init() {

	this.__initRootBasedExpressionDelegation()
}

func (this Iterator[Chain_T, Node_T]) SetIterationChain(chain Chain_T) {

	this.root_iterator_delegator.SetLeftOperand(chain)
}

func (this *Iterator[Chain_T, Node_T]) GetIterationChain() Chain_T {

	return this.root_iterator_delegator.GetLeftOperand()
}

func (this *Iterator[Chain_T, Node_T]) __initRootBasedExpressionDelegation() {

	this.root_iterator_delegator.Accept(this.root.GetBasedExpression())
}

func (this Iterator[Chain_T, Node_T]) Root() *operator.SecondaryOperator {

	return this.root
}

func (this *Iterator[Chain_T, Node_T]) SetRoot(root *operator.SecondaryOperator) {

	this.root = root
}
