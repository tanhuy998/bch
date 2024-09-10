package memoryCache

import "sync"

type (
	cache_unit[Key_T, Value_T any] struct {
		sync.Map
	}

	//cache_value = sync.Map

	CommitFunction[Value_T any] func(val Value_T)
	ReadUnlockFunction          func()
)

/*
Return the copy of the cached value
this method does not lock the cache value
*/
func (this *cache_unit[Key_T, Value_T]) ReadInstanctly(key Key_T) (value Value_T, exists bool) {

	cache, exists := this.getCache(key)

	if !exists {

		return
	}

	return cache.value, exists
}

/*
Return the copy of the cached value and lock the cache value
until the release lock func invoked,
if the cached value is not present
releaseLock func will be nil, ensure check of the second return value for
the existence of the cached value
*/
func (this *cache_unit[Key_T, Value_T]) ReadLock(key Key_T) (value Value_T, exists bool, releaseLock ReadUnlockFunction) {

	cache, exists := this.getCache(key)

	if !exists {

		return
	}

	cache.RLock()

	releaseLock = func() {

		cache.RUnlock()
	}

	return cache.value, true, releaseLock
}

func (this *cache_unit[Key_T, Value_T]) getCache(key Key_T) (cache *cache_value[Value_T], exists bool) {

	unknown, exists := this.Load(key)

	if !exists {

		return
	}

	if c, ok := unknown.(*cache_value[Value_T]); ok {

		return c, true
	}

	return
}

/*
Instantly Create a room for the value that corresponding to the key
if the cached key exists, the progress wiil lock for write operations
*/
func (this *cache_unit[Key_T, Value_T]) Set(key Key_T, value Value_T) {

	cache, exists := this.getCache(key)

	if !exists {

		cache = new(cache_value[Value_T])
		return
	}

	cache.Lock()
	defer cache.Unlock()

	cache.value = value
}

/*
Instantly Delete the room of the cached value corresponding to the given key,
The delete progress locks on both read and write operations
*/
func (this *cache_unit[Key_T, Value_T]) Delete(key Key_T) (deleted bool) {

	cache, exists := this.getCache(key)

	if !exists {

		return
	}

	cache.Lock()
	cache.RLock()
	defer cache.RUnlock()
	defer cache.Unlock()

	this.Map.Delete(key)
	return true
}

/*
Update a cached value, when the update method invoked, its locks until
the commitFunc commits the value to the cache room
*/
func (this *cache_unit[Key_T, Value_T]) Update(key Key_T) (value Value_T, commitFunc CommitFunction[Value_T]) {

	cache, exists := this.getCache(key)

	if !exists {

		return
	}

	cache.Lock()

	commitFunc = func(val Value_T) {

		cache.value = val
		cache.Unlock()
	}

	return cache.value, commitFunc
}
