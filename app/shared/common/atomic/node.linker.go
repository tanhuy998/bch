package atomic

import "app/shared/common/constraint"

type (
	INodeValueManipulator[T any] interface {
		Value() T
		SetValue(v T)
	}

	IListNode[Actual_Node_T any] interface {
		constraint.IMutex
		constraint.IRMutex
		INodeActualType[Actual_Node_T]
		isNode()
	}

	INodeActualType[T any] interface {
		self() T
	}

	INextNodeReference[T any] interface {
		IListNode[T]
		getNext() INextNodeReference[T]
		setNext(node INextNodeReference[T])
	}

	IPreviousNodeReference[T any] interface {
		IListNode[T]
		getPrev() IPreviousNodeReference[T]
		setPrev(node IPreviousNodeReference[T])
	}

	IDoubleNodeManipulator[T any] interface {
		INextNodeReference[T]
		IPreviousNodeReference[T]
	}
)

func linkBefore[Node_T IDoubleNodeManipulator[Node_T]](
	node IDoubleNodeManipulator[Node_T], unlinked IDoubleNodeManipulator[Node_T],
) {

	switch {
	case node == nil:
		panic("")
	case unlinked == nil:
		return
	}

	var temp IDoubleNodeManipulator[Node_T] = node.getPrev().self()

	// node.prev = unlinked
	// unlinked.next = node
	// unlinked.prev = temp

	node.setPrev(unlinked)
	unlinked.setNext(node)
	unlinked.setPrev(temp)
}

func linkAfter[Node_T IDoubleNodeManipulator[Node_T]](
	node IDoubleNodeManipulator[Node_T], unlinked IDoubleNodeManipulator[Node_T],
) {

	switch {
	case node == nil:
		panic("")
	case unlinked == nil:
		return
	}

	var temp IDoubleNodeManipulator[Node_T] = node.getNext().self()

	// node.next = unlinked
	// unlinked.prev = node
	// unlinked.next = temp

	node.setNext(unlinked)
	unlinked.setPrev(node)
	unlinked.setNext(temp)
}

func swapLink[Node_T IDoubleNodeManipulator[Node_T]](
	left IDoubleNodeManipulator[Node_T], right IDoubleNodeManipulator[Node_T],
) {

}
