package atomic

type (
	ILinkedListIteratorNodeValueManipulator[T any] interface {
		Value() (T, bool)
		SetValue(v T) bool
	}

	ILinkedListIterator[Value_T any] interface {
		ILinkedListIteratorNodeValueManipulator[Value_T]
		Next() bool
	}

	ILinkedListSeeker[Value any] interface {
		ILinkedListIterator[Value]
		Previous() bool
	}
)

type (
	/*
		Iterator is object that could iterate through a linked list's nodes
		once an iterator is distrubted by the list, the linked list do the read lo
	*/
	double_iterator_t[Value_T any] struct {
		//*node_synchronizer_t
		*SynchronizeEmitter
		pending *node_double_t[Value_T]
		loaded  *node_double_t[Value_T]
	}
)

func newDoubleLinkedListIterator[Value_T any](l *seekable_double_linked_list_t[Value_T]) ILinkedListIterator[Value_T] {

	return &double_iterator_t[Value_T]{
		//SynchronizeEmitter: &l.synchronizer,
		loaded: l.head,
	}
}

/*
Next() iterates next linked node of the list chain.
*/
func (this *double_iterator_t[Value_T]) Next() bool {

	this.pending.RUnlock()
	this.SynchronizeEmitter.Lock()
	defer this.SynchronizeEmitter.Unlock()

	switch {
	case !this.iterateActualNext():
		return false
	default:
		this.pending.RLock()
		return true
	}
}

func (this *double_iterator_t[Value_T]) TryNext() (nextAccquired bool, stillHolds bool) {

	return false, false
}

func (this *double_iterator_t[Value_T]) UnBind() {

	this.pending = nil
	this.loaded = nil
	this.SynchronizeEmitter = nil
}

func (this *double_iterator_t[Value_T]) iterateActualNext() bool {

	this.pending.Sync()

	for pending, loaded := this.pending, this.loaded; pending != nil; this.pending, this.loaded = loaded, pending.next {

		switch {
		case !pending.linked && loaded != nil:
			continue
		case pending.linked:
			return true
		}
	}

	return false
}

func (this *double_iterator_t[Value_T]) Value() (v Value_T, ok bool) {

	switch {
	case this.pending == nil:
		ok = false
		return
	default:
		return this.pending.val, true
	}
}

func (this *double_iterator_t[Value_T]) SetValue(v Value_T) bool {

	switch {
	case this.pending == nil:
		return false
	default:
		this.pending.RUnlock()
		defer this.pending.Lock()

		this.pending.SetValue(v)

		return true
	}
}
