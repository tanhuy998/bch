package memoryCache

import "sync"

type (
	cache_value[Value_T any] struct {
		sync.RWMutex
		value Value_T
		//isLock bool
	}
)

// func (this *cache_value[Value_T]) Lock() {

// 	if this.isLock {

// 		return
// 	}

// 	this.isLock = true
// 	this.RWMutex.Lock()
// }

// func (this *cache_value[Value_T]) UnLock() {

// 	if !this.isLock {

// 		return
// 	}

// 	this.isLock = false
// 	this.RWMutex.Unlock()
// }

func (this *cache_value[Value_T]) GetValue() Value_T {

	return this.value
}
