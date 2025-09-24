package libCommon

import "sync"

type (
	WriteConcernMap[Key_T comparable, Value_T comparable] struct {
		sync.RWMutex
		m map[Key_T]Value_T
	}
)

func (this *WriteConcernMap[Key_T, Value_T]) CompareAndDelete(key Key_T, old Value_T) (deleted bool) {

	this.Lock()
	defer this.Unlock()

	switch v, existed := this.m[key]; {
	case !existed:
		return false
	case v == old:
		delete(this.m, key)
		return true
	default:
		return false
	}
}

func (this *WriteConcernMap[Key_T, Value_T]) CompareAndSwap(key Key_T, old Value_T, new Value_T) bool {

	this.Lock()
	defer this.Unlock()

	switch v, existed := this.m[key]; {
	case !existed:
		return false
	case v == old:
		this.m[key] = new
		return true
	default:
		return false
	}
}

func (this *WriteConcernMap[Key_T, Value_T]) Delete(Key Key_T) {

	this.Lock()
	defer this.Unlock()

	delete(this.m, Key)
}

func (this *WriteConcernMap[Key_T, Value_T]) Load(key Key_T) (val Value_T, ok bool) {

	this.RLock()
	defer this.RUnlock()

	val, ok = this.m[key]
	return
}

func (this *WriteConcernMap[Key_T, Value_T]) LoadAndDelete(key Key_T) (val Value_T, deleted bool) {

	this.Lock()
	defer this.Unlock()

	val, existed := this.m[key]

	if existed {

		delete(this.m, key)
		return val, true
	}

	deleted = false
	return
}

func (this *WriteConcernMap[Key_T, Value_T]) LoadOrStore(key Key_T, val Value_T) (actual Value_T, loaded bool) {

	this.Lock()
	defer this.Unlock()

	v, existed := this.m[key]

	if existed {

		return v, true
	}

	this.m[key] = val

	return val, false
}

func (this *WriteConcernMap[Key_T, Value_T]) Range(
	fn func(key Key_T, val Value_T) bool,
) {

	if fn == nil {

		return
	}

	this.Lock()
	defer this.Unlock()

	for k, v := range this.m {

		fn(k, v)
	}
}

func (this *WriteConcernMap[Key_T, Value_T]) Store(key Key_T, val Value_T) {

	this.Lock()
	defer this.Unlock()

	this.m[key] = val
}

func (this *WriteConcernMap[Key_T, Value_T]) Swap(key Key_T, val Value_T) (previous Value_T, loaded bool) {

	this.Lock()
	defer this.__set(key, val)
	defer this.Unlock()

	switch v, existed := this.m[key]; {
	case existed:
		return v, true
	default:
		return
	}
}

// lock free method
func (this *WriteConcernMap[Key_T, Value_T]) __set(key Key_T, val Value_T) {

	this.m[key] = val
}
