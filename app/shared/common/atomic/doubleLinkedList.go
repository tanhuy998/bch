package atomic

type (
	DoubleLinkedListFloatingPolicy[Value_T any] func(currentVal Value_T, comparedValue Value_T) (shouldSwap bool)
)

type (
	// mutex based double linked list implementation
	// for linked list atomicy synchronization in concurrent enviroment
	DoubleLinkedList[Value_T any] struct {
		seekable_double_linked_list_t[Value_T]
		floatPolicy DoubleLinkedListFloatingPolicy[Value_T]
	}
)

func (this *DoubleLinkedList[Value_T]) swap(l *node_double_t[Value_T], r *node_double_t[Value_T]) {

}

func (this *DoubleLinkedList[Value_T]) Unlink(
	fromIndex int, toIndex int,
) (headOfUnlinked *node_double_t[Value_T], tailOfUnlinked *node_double_t[Value_T]) {

	switch {
	case toIndex < fromIndex:
		panic("")
	case toIndex > this.length:
		panic("")
	case fromIndex < 0:
		panic("")
	}

	sync := this.synchronizer.LockAndSyncLater()

	defer func() {

		sync(this.length)
	}()

	headOfUnlinked = this.nodeAt(fromIndex)
	tailOfUnlinked = this.nodeAt(toIndex)

	remainTail := headOfUnlinked.next
	remainHead := tailOfUnlinked.prev

	if remainTail != nil {
		remainTail.next = remainHead
	} else {
		this.head = remainTail
	}

	if remainHead != nil {
		remainHead.prev = remainTail
	} else {
		this.tail = remainHead
	}

	iterator := headOfUnlinked

	for ; iterator != nil && iterator != tailOfUnlinked; iterator = iterator.next {
		iterator.linked = false
	}

	return
}

func (this *DoubleLinkedList[Value_T]) PopBack() {

	switch {
	case this.length == 0:
		return
	default:
		this.Unlink(this.length-1, this.length-1)
	}
}

func (this *DoubleLinkedList[Value_T]) PopFront() {
	switch {
	case this.length == 0:
		return
	default:
		this.Unlink(0, 0)
	}
}

func (this *DoubleLinkedList[Value_T]) sync() {

	this.synchronizer.Sync(this.length)
}

func (this *DoubleLinkedList[Value_T]) AppendAt(index int, vals ...Value_T) {

	this.synchronizer.Lock()
	defer this.sync()
	defer this.synchronizer.Unlock()

	pivot := this.nodeAt(index)

	if pivot == nil {
		this.synchronizer.Unlock()
		panic("")
	}

	appendHead, appendTail := this.enchain(vals...)

	nodeAfterPivot := pivot.next

	pivot.next = appendHead

	switch {
	case nodeAfterPivot == nil:
		this.tail = appendTail
	default:
		nodeAfterPivot.next = appendTail
	}

	this.length += len(vals)
}

func (this *DoubleLinkedList[Value_T]) PrependAt(index int, vals ...Value_T) {

	if index+1 > this.length {
		panic("prepend at index out of range")
	}

	if index > 0 {
		this.AppendAt(index - 1)
		return
	}

	this.synchronizer.Lock()
	defer this.sync()
	defer this.synchronizer.Unlock()

	// index == 0
	appendHead, appendTail := this.enchain(vals...)

	lastListHead := this.head
	this.head = appendHead
	appendTail.next = lastListHead

	this.length += len(vals)
}

func (this *DoubleLinkedList[Value_T]) linkTwoNode(left *node_double_t[Value_T], right *node_double_t[Value_T]) {

	left.next = right
	right.prev = left
}

func (this *DoubleLinkedList[Value_T]) link(nodes ...*node_double_t[Value_T]) {

	for i := range len(nodes) {

		switch {
		case i+1 == len(nodes):
			return
		default:
			this.linkTwoNode(nodes[i], nodes[i+1])
		}
	}
}

func (this *DoubleLinkedList[Value_T]) enchain(values ...Value_T) (head *node_double_t[Value_T], tail *node_double_t[Value_T]) {

	if len(values) == 0 {
		panic("enchaining empty slice of value for double linked list")
	}

	var iterator **node_double_t[Value_T] = &head

	for _, v := range values {

		newNode := this.distributeNewNode()
		newNode.val = v

		*iterator = newNode
		iterator = &(*iterator).next
	}

	tail = (*iterator).prev

	return
}

func (this *DoubleLinkedList[Value_T]) PushBack(v Value_T) {

	this.AppendAt(this.length-1, v)
}

func (this *DoubleLinkedList[Value_T]) PushFront(v Value_T) {

	this.PrependAt(0, v)
}

func (this *DoubleLinkedList[Value_T]) distributeNewNode() *node_double_t[Value_T] {

	ret := new(node_double_t[Value_T])

	ret.linked = true
	ret.synchronizer = this.synchronizer.Distribute()

	return ret
}
