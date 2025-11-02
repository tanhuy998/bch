package atomic

type (
	seekable_double_linked_list_t[Value_T any] struct {
		synchronizer SynchronizeEmitter // linked_list_synchronizer_t
		length       int
		head         *node_double_t[Value_T]
		tail         *node_double_t[Value_T]
	}
)

func (this *seekable_double_linked_list_t[Value_T]) ValueAt(index int) (v Value_T, ok bool) {

	this.synchronizer.RLock()
	defer this.synchronizer.RUnlock()

	node := this.nodeAt(index)

	switch {
	case node == nil:
		ok = false
		return
	default:
		node.RLock()
		defer node.RUnlock()
		v = node.val
		ok = true
		return
	}
}

func (this *seekable_double_linked_list_t[Value_T]) Interate(from int) ILinkedListIterator[Value_T] {
	return newDoubleLinkedListIterator(this)
}

func (this *seekable_double_linked_list_t[Value_T]) ValueAtHead() (v Value_T, ok bool) {
	this.synchronizer.RLock()
	defer this.synchronizer.RUnlock()

	switch {
	case this.head == nil:
		return
	default:
		return this.head.val, true
	}
}

func (this *seekable_double_linked_list_t[Value_T]) Head() *node_double_t[Value_T] {

	return this.head
}

func (this *seekable_double_linked_list_t[Value_T]) Tail() *node_double_t[Value_T] {

	return this.tail
}

func (this *seekable_double_linked_list_t[Value_T]) ValueAtTail() (v Value_T, ok bool) {
	this.synchronizer.RLock()
	defer this.synchronizer.RUnlock()

	switch {
	case this.tail == nil:
		return
	default:
		return this.tail.val, true
	}
}

func (this *seekable_double_linked_list_t[Value_T]) nodeAt(index int) *node_double_t[Value_T] {

	if index+1 > this.length {
		return nil
	}

	current := this.head

	for i := 0; i <= index; i++ {

		if current.next == nil {
			break
		}

		current = current.next
	}

	return current
}
