package atomic

type (
	node_double_t[Value_T any] struct {
		abstract_node_t[Value_T]
		next *node_double_t[Value_T]
		prev *node_double_t[Value_T]
	}
)

func (this *node_double_t[Value_T]) getNext() INextNodeReference[*node_double_t[Value_T]] {

	return this.next
}

func (this *node_double_t[Value_T]) setNext(node INextNodeReference[*node_double_t[Value_T]]) {

	switch {
	case node != nil:
		this.next = nil
	default:
		this.next = node.self()
	}
}

func (this *node_double_t[Value_T]) self() *node_double_t[Value_T] {

	return this
}

func (this *node_double_t[Value_T]) getPrev() IPreviousNodeReference[*node_double_t[Value_T]] {

	return this.prev
}

func (this *node_double_t[Value_T]) setPrev(node IPreviousNodeReference[*node_double_t[Value_T]]) {

	switch {
	case node == nil:
		this.prev = nil
	default:
		this.prev = node.self()
	}
}
