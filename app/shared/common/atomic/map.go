package atomic

import "sync"

type (
	Map[Key_T comparable, Value_T any] struct {
		mutex sync.RWMutex
		m     map[Key_T]Value_T
	}
)

func (this *Map[Key_T, Value_T]) init() {

	if this.m == nil {
		this.m = make(map[Key_T]Value_T)
	}
}

func (this *Map[Key_T, Value_T]) Load(key Key_T) (v Value_T, ok bool) {

	this.mutex.RLock()
	defer this.mutex.RUnlock()

	if len(this.m) == 0 {
		return
	}

	v, ok = this.m[key]
	return
}

func (this *Map[Key_T, Value_T]) Delete(key Key_T) (lastValue Value_T, ok bool) {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	if len(this.m) == 0 {
		return
	}

	lastValue, ok = this.m[key]

	delete(this.m, key)
	return
}

func (this *Map[Key_T, Value_T]) Set(key Key_T, val Value_T) (lastValue Value_T, ok bool) {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.init()

	lastValue, ok = this.m[key]

	this.m[key] = val
	return
}

func (this *Map[Key_T, Value_T]) SetIfNotExist(key Key_T, val Value_T) (ok bool) {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.init()

	_, existed := this.m[key]

	switch {
	case existed:
		return false
	default:
		this.m[key] = val
		return true
	}
}

func (this *Map[Key_T, Value_T]) LoadOrStore(key Key_T, val Value_T) (actual Value_T, loaded bool) {

	this.mutex.Lock()
	defer this.mutex.Unlock()

	v, loaded := this.m[key]

	switch {
	case loaded:
		actual = v
		return
	default:
		this.m[key] = val
		return val, false
	}
}
