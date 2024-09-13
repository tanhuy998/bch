package memoryCache

import (
	"sync"
	"time"
)

type (
	ICacheRoom[Value_T any] interface {
		GetValue() Value_T
		SetValue(val Value_T)
	}

	IExpirationCacheRoom[Value_T any] interface {
		ICacheRoom[Value_T]
		GetTimer() *time.Timer
		GetExpirationTime() *time.Time
	}

	/*
		cache_value represents the a key room for holding the cache value
	*/
	cache_value[Value_T any] struct {
		sync.RWMutex
		value       Value_T
		expireTimer *time.Timer
		/*
			cache_value.deactivated is a state that indicates the current cache
			room is unlink with the cache_unit map (in order words, a key is deleted from the cache_unit)
			goroutines that hold this room for write operations (such as update and reassign cache room value)
			but have to wait for another to unlock write (the situation is the other goroutine delete the key corresponding to this room).
		*/
		deactivated bool
	}

	// expire_cache_value[Value_T any] struct {
	// 	cache_value[Value_T]
	// 	moment time.Time
	// }
)

func deactivateCacheValue[Value_T any](cache *cache_value[Value_T]) {

	if cache.expireTimer == nil {

		return
	}

	cache.expireTimer.Stop()
	cache.expireTimer = nil
}

// func (this *cache_value[Value_T]) GetValue() Value_T {

// 	return this.value
// }
// func (this *cache_value[Value_T]) SetValue(val Value_T) {

// 	this.value = val
// }

// func (this *expire_cache_value[Value_T]) GetValue() Value_T {

// 	return this.value
// }

// func (this *expire_cache_value[Value_T]) SetValue(val Value_T) {

// 	this.value = val
// }

// func (this *expire_cache_value[Value_T]) GetTimer() *time.Timer {

// 	return nil
// }

// func (this *expire_cache_value[Value_T]) GetExpirationTime() *time.Time {

// 	return &this.moment
// }

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
