package memoryCache

import (
	libTryLock "app/src/internal/lib/tryLock"
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ERR_INVALID_VALUE_TYPE   = errors.New("cache unit error: invalid value type, cannot resolve cache value")
	ERR_EXPIRATION_MOMENT    = errors.New("cache expiration error: invalid expire time")
	ERR_OUT_OF_CONTEXT       = errors.New("cache err: context done before reaching operation")
	ERR_OBSOLETE_CACHE_VALUE = errors.New("cache err: obsolete cache value, try the operation again")
	ERR_LOCKING_SAME_KEY     = errors.New("cache err: locking the same key")
)

type (
	ICacheUnit[Key_T, Value_T comparable] interface {
		ReadInstanctly(ctx context.Context, key Key_T) (value Value_T, exists bool)
		ReadAndHold(ctx context.Context, key Key_T) (value Value_T, exists bool, releaseLock ReadUnlockFunction)
		Set(ctx context.Context, key Key_T, value Value_T)
		Delete(ctx context.Context, key Key_T) (deleted bool)
		Update(ctx context.Context, key Key_T) (value Value_T, keyExists bool, commitFunc CommitFunction[Value_T])
	}

	ISelfCleanupCacheUnit interface {
		cleanup()
	}

	CacheUnitCleanupFunction[Key_T, Value_T comparable] func(key interface{}, value interface{}) bool

	cache_unit[Key_T, Value_T comparable] struct {
		sync.Map

		//cleanupFunc CacheUnitCleanupFunction[Key_T, Value_T]
	}

	//cache_value = sync.Map

	CommitFunction[Value_T comparable] func(val Value_T)
	RevokeUpdateFunction               func()
	ReadUnlockFunction                 func()

	UpdateCommandFunction[Value_T comparable] func() (CommitFunction[Value_T], RevokeUpdateFunction)
	//HoldCommandFunction                func() (libTryLock.ContextDoneCheckFunction, ReadUnlockFunction)

	CommitFunctionOption                       int
	CommitFunctionOptionActivator[Value_T any] func(cache *cache_value[Value_T], oldVal Value_T, newVal Value_T)
)

func newCacheUnit[Key_T, Value_T comparable]() *cache_unit[Key_T, Value_T] {

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
func (this *cache_unit[Key_T, Value_T]) Read(
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

	lockAccquired, release, err := libTryLock.AccquireReadLock(ctx, &cache.RWMutex)

	if err != nil {

		return
	}

	if lockAccquired {

		defer release()
	}

	if ctx.Err() != nil {

		err = ERR_OUT_OF_CONTEXT
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
func (this *cache_unit[Key_T, Value_T]) Hold(
	ctx context.Context, key Key_T,
) (value Value_T, exists bool, release func(), err error) {

	if this.isKeyConflict(ctx, key) {

		err = ERR_LOCKING_SAME_KEY
		return
	}

	cache, exists, e := this.getCache(ctx, key)

	if e != nil {

		err = e
		return
	}

	if !exists {

		return
	}

	lockAccquired, releasAcquiredLock, err := libTryLock.AccquireReadLock(ctx, cache)

	if err != nil {

		return
	}

	if lockAccquired {

		release = func() {

			releasAcquiredLock()
		}
	}

	if ctx.Err() != nil {

		if lockAccquired {
			releasAcquiredLock()
		}

		err = ERR_OUT_OF_CONTEXT
		return
	}

	return cache.value, true, release, nil
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
func (this *cache_unit[Key_T, Value_T]) Set(ctx context.Context, key Key_T, value Value_T) (returnErr error) {

	cache, exists, err := this.getCache(ctx, key)

	if err != nil {

		return err
	}

	if !exists || cache.deactivated {

		cache = new(cache_value[Value_T])

		defer func() {

			if returnErr != nil {

				this.Map.Swap(key, cache)
			}
		}()
	}

	return this.set(ctx, cache, value)
}

func (this *cache_unit[Key_T, Value_T]) set(ctx context.Context, cache *cache_value[Value_T], value Value_T) error {

	if cache.deactivated {

		return ERR_OBSOLETE_CACHE_VALUE
	}

	lockAcquired, release, err := libTryLock.AcquireLock(ctx, cache)

	if err != nil {

		return err
	}

	if ctx.Err() != nil {

		return ERR_OUT_OF_CONTEXT
	}

	if !lockAcquired {

		return ERR_OUT_OF_CONTEXT
	}

	defer release()

	cache.value = value
	return nil
}

/*
Instantly Create expiration room for the value that corresponding to the key
if the cache key exist, the key will be override with new value and
*/
func (this *cache_unit[Key_T, Value_T]) SetWithExpire(
	ctx context.Context, key Key_T, value Value_T, moment time.Time,
) (err error) {

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

	_, command, err := this.startModify(ctx, currentCache)

	if err != nil {

		return err
	}

	commit, revoke := command()

	if ctx.Err() != nil {

		revoke()
		err = ERR_OUT_OF_CONTEXT
		return
	}

	cleanupTime := time.AfterFunc(duration, func() {

		deactivateCacheValue(currentCache)
		this.Delete(ctx, key)
	})

	currentCache.expireTimer = cleanupTime
	commit(value)

	return nil
}

/*
Instantly Delete the room of the cached value corresponding to the given key,
The delete progress locks on both read and write operations
*/
func (this *cache_unit[Key_T, Value_T]) Delete(ctx context.Context, key Key_T) (deleted bool, err error) {

	cache, exists, err := this.getCache(ctx, key)

	if err != nil {

		return
	}

	if !exists {

		return
	}

	_, revoke, err := this.startModify(ctx, cache)
	defer revoke()

	if err != nil {

		return
	}

	cache.deactivated = false

	if cache.expireTimer != nil {

		deactivateCacheValue(cache)
	}

	this.Map.Delete(key)
	return true, nil
}

/*
Modify a cached value, when the update method invoked, its locks until
the commitFunc commits the value to the cache room
*/
func (this *cache_unit[Key_T, Value_T]) Modify(
	ctx context.Context, key Key_T,
) (value Value_T, keyExists bool, updateCommand UpdateCommandFunction[Value_T], err error) {

	if this.isKeyConflict(ctx, key) {

		err = ERR_LOCKING_SAME_KEY
		return
	}

	cache, keyExists, err := this.getCache(ctx, key)

	if err != nil || !keyExists {

		return
	}

	value, updateCommand, err = this.startModify(ctx, cache)

	if err == ERR_OBSOLETE_CACHE_VALUE {

		keyExists = false
		err = nil
	}

	return
}

func (this *cache_unit[Key_T, Value_T]) startModify(
	ctx context.Context, cache *cache_value[Value_T],
) (value Value_T, updateCommand UpdateCommandFunction[Value_T], err error) {

	if cache.deactivated {

		err = ERR_OBSOLETE_CACHE_VALUE
		return
	}

	readLockAcquired, releaseRead, err := libTryLock.AccquireReadLock(ctx, cache)

	if err != nil {

		return
	}

	writeLockAcquired, releaseWrite, err := libTryLock.AcquireLock(ctx, cache)

	if err != nil {

		return
	}

	if readLockAcquired && writeLockAcquired {

		updateCommand = func() (commit CommitFunction[Value_T], abort RevokeUpdateFunction) {

			commit = func(val Value_T) {

				cache.value = val
				releaseRead()
				releaseWrite()
			}

			abort = func() {

				cache.Unlock()
				releaseRead()
				releaseWrite()
			}

			return
		}

		return cache.value, updateCommand, nil
	}

	if readLockAcquired {
		releaseRead()
	}

	if writeLockAcquired {
		releaseWrite()
	}

	err = ERR_OUT_OF_CONTEXT
	return
}

func (this *cache_unit[Key_T, Value_T]) isKeyConflict(ctx any, key Key_T) bool {

	if !isKeyLockingContext[Key_T](ctx) {

		return false

	}

	v, ok := ctx.(IKeyLockingContext[Key_T])

	return ok && v.HasKey(key)
}
