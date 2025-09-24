package stack

import "sync"

type (
	stack_node_t[T any] struct {
		val  T
		prev *stack_node_t[T]
	}

	Stack[Val_T any] struct {
		sync.Mutex
		head *stack_node_t[Val_T]
		mass int
	}
)

func New[Val_T any]() *Stack[Val_T] {

	return new(Stack[Val_T])
}

func (this *Stack[Val_T]) IsEmpty() bool {

	return this.head == nil
}

func (this *Stack[Val_T]) Mass() int {

	return this.mass
}

func (this *Stack[Val_T]) Push(val Val_T) {

	this.Lock()
	defer this.Unlock()

	this._push(val)
}

func (this *Stack[Val_T]) _push(val Val_T) {

	defer func() {

		this.mass++
	}()

	cur := this.head

	newHead := new(stack_node_t[Val_T])
	newHead.val = val

	this.head = newHead

	switch {
	case cur == nil:
		return
	default:
		this.head.prev = cur
	}
}

func (this *Stack[Val_T]) Pop() (val Val_T, ok bool) {

	this.Lock()
	defer this.Unlock()

	return this._pop()
}

func (this *Stack[Val_T]) _pop() (val Val_T, ok bool) {

	if this.head == nil {

		return
	}

	val = this.head.val
	ok = true

	popped_head := this.head

	this.head = this.head.prev
	popped_head.prev = nil
	this.mass--

	return
}

func (this *Stack[Val_T]) Cleanup() {

	this.Lock()
	defer this.Unlock()

	this.mass = 0

	for this.head != nil {

		popped_head := this.head

		this.head = this.head.prev

		popped_head.prev = nil
	}
}

func (this *Stack[Val_T]) Snapshot() []Val_T {

	this.Lock()
	defer this.Unlock()

	return this._snapShot()
}

func (this *Stack[Val_T]) _snapShot() []Val_T {

	ret := make([]Val_T, this.mass)

	var iterator *stack_node_t[Val_T] = this.head

	for i := range ret {

		headIndex := this.mass - 1 - i

		ret[headIndex] = iterator.val

		iterator = iterator.prev
	}

	return ret
}

func (this *Stack[Val_T]) Clone() *Stack[Val_T] {

	this.Lock()
	defer this.Unlock()

	ret := new(Stack[Val_T])

	snapShot := this._snapShot()

	for _, val := range snapShot {

		ret._push(val)
	}

	return ret
}
