package memoryCache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ERR_INVALID_VALUE_TYPE = errors.New("cache unit error: invalid value type, cannot resolve cache value")
	ERR_EXPIRATION_MOMENT  = errors.New("cache expiration error: invalid expire time")
)

type (
	ICacheUnit[Key_T, Value_T any] interface {
		ReadInstanctly(ctx context.Context, key Key_T) (value Value_T, exists bool)
		ReadAndHold(ctx context.Context, key Key_T) (value Value_T, exists bool, releaseLock ReadUnlockFunction)
		Set(ctx context.Context, key Key_T, value Value_T)
		Delete(ctx context.Context, key Key_T) (deleted bool)
		Update(ctx context.Context, key Key_T) (value Value_T, keyExists bool, commitFunc CommitFunction[Value_T])
	}

	ISelfCleanupCacheUnit interface {
		cleanup()
	}

	CacheUnitCleanupFunction[Key_T, Value_T any] func(key interface{}, value interface{}) bool

	cache_unit[Key_T, Value_T any] struct {
		sync.Map

		//cleanupFunc CacheUnitCleanupFunction[Key_T, Value_T]
	}

	//cache_value = sync.Map

	CommitFunction[Value_T any] func(val Value_T)
	RevokeUpdateFunction        func()
	ReadUnlockFunction          func()

	UpdateCommandFunction[Value_T any] func() (CommitFunction[Value_T], RevokeUpdateFunction)

	CommitFunctionOption                       int
	CommitFunctionOptionActivator[Value_T any] func(cache *cache_value[Value_T], oldVal Value_T, newVal Value_T)
)

func newCacheUnit[Key_T, Value_T any]() *cache_unit[Key_T, Value_T] {

	cacheUnit := new(cache_unit[Key_T, Value_T])

	// cacheUnit.cleanupFunc = func(key interface{}, value interface{}) bool {

	// 	var (
	// 		actualKey  Key_T
	// 		cacheValue *cache_value[Value_T]
	// 	)

	// 	if v, ok := key.(Key_T); ok {

	// 		actualKey = v
	// 	} else {

	// 		return true
	// 	}

	// 	if v, ok := value.(*cache_value[Value_T]); ok {

	// 		cacheValue = v
	// 	} else {

	// 		return true
	// 	}

	// 	if time.Now().Before(cacheValue.expireTime) ||
	// 		cacheValue.expireTime.IsZero() {

	// 		return true
	// 	}

	// 	cacheUnit.Delete(actualKey)

	// 	return true
	// }

	return cacheUnit
}

// func (this *cache_unit[Key_T, Value_T]) cleanup() {

// 	this.Map.Range(this.cleanupFunc)
// }

/*
Return the copy of the cached value
this method does not lock the cache value
*/
func (this *cache_unit[Key_T, Value_T]) ReadInstanctly(
	ctx context.Context, key Key_T,
) (value Value_T, exists bool, err error) {

	cache, exists, e := this.getCache(ctx, key)

	if e != nil {

		err = e
		return
	}

	if !exists {

		return
	}

	return cache.value, exists, nil
}

/*
Return the copy of the cached value and lock the cache value
until the release lock func invoked,
if the cached value is not present
releaseLock func will be nil, ensure check of the second return value for
the existence of the cached value
*/
func (this *cache_unit[Key_T, Value_T]) ReadAndHold(
	ctx context.Context, key Key_T,
) (value Value_T, exists bool, releaseLock ReadUnlockFunction, err error) {

	cache, exists, e := this.getCache(ctx, key)

	if e != nil {

		err = e
		return
	}

	if !exists {

		return
	}

	cache.RLock()

	releaseLock = func() {

		cache.RUnlock()
	}

	return cache.value, true, releaseLock, nil
}

func (this *cache_unit[Key_T, Value_T]) getCache(
	ctx context.Context, key Key_T,
) (cache *cache_value[Value_T], exists bool, err error) {

	unknown, exists := this.Map.Load(key)

	if !exists {

		return
	}

	if cache, ok := unknown.(*cache_value[Value_T]); ok && !cache.deactivated {

		return cache, true, nil

	} else if !cache.deactivated {

		return nil, false, nil
	}

	err = ERR_INVALID_VALUE_TYPE
	return
}

/*
Instantly Create a room for the value that corresponding to the key
if the cached key exists, the progress wiil lock for write operations.
This method is an idempotent operation
*/
func (this *cache_unit[Key_T, Value_T]) Set(ctx context.Context, key Key_T, value Value_T) error {

	cache, exists, err := this.getCache(ctx, key)

	if err != nil {

		return err
	}

	if !exists || cache.deactivated {

		cache = new(cache_value[Value_T])
		this.Map.Swap(key, cache)
	}

	cache.Lock()
	defer cache.Unlock()

	cache.value = value
	return nil
}

/*
Instantly Create expiration room for the value that corresponding to the key
if the cache key exist, the key will be override with new value and
*/
func (this *cache_unit[Key_T, Value_T]) SetWithExpire(
	ctx context.Context, key Key_T, value Value_T, moment time.Time,
) error {

	duration := time.Until(moment)

	if duration <= 0 {
		/*
			similar expression but a bit slower
			moment.Before(now) || moment.Equal(now)
		*/
		return ERR_EXPIRATION_MOMENT
	}

	currentCache, exists, err := this.getCache(ctx, key)

	if err != nil {

		return err
	}

	if !exists {

		currentCache = new(cache_value[Value_T])
		this.Map.Store(key, currentCache)
	}

	currentCache.value = value

	cleanupTime := time.AfterFunc(duration, func() {

		deactivateCacheValue(currentCache)
		this.Delete(ctx, key)
	})

	currentCache.expireTimer = cleanupTime

	return nil
}

/*
Instantly Delete the room of the cached value corresponding to the given key,
The delete progress locks on both read and write operations
*/
func (this *cache_unit[Key_T, Value_T]) Delete(ctx context.Context, key Key_T) (deleted bool) {

	cache, exists, _ := this.getCache(ctx, key)

	if !exists {

		return
	}

	cache.Lock()
	cache.RLock()

	defer cache.RUnlock()
	defer cache.Unlock()

	cache.deactivated = false

	if cache.expireTimer != nil {

		deactivateCacheValue(cache)
	}

	this.Map.Delete(key)
	return true
}

/*
Update a cached value, when the update method invoked, its locks until
the commitFunc commits the value to the cache room
*/
func (this *cache_unit[Key_T, Value_T]) Update(
	ctx context.Context, key Key_T,
) (value Value_T, keyExists bool, updateCommand UpdateCommandFunction[Value_T], err error) {

	cache, keyExists, err := this.getCache(ctx, key)

	if err != nil || !keyExists {

		return
	}

	if cache.deactivated {

		keyExists = false
		return
	}

	cache.Lock()

	updateCommand = func() (commit CommitFunction[Value_T], revoke RevokeUpdateFunction) {

		commit = func(val Value_T) {

			cache.value = val
			cache.Unlock()
		}

		revoke = func() {

			cache.Unlock()
		}

		return
	}

	return cache.value, true, updateCommand, nil
}
