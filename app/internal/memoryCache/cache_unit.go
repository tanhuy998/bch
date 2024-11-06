package memoryCache

import (
	libTryLock "app/internal/lib/tryLock"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
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
		topic string
		*log.Logger
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

func newCacheUnit[Key_T, Value_T comparable](topic string) *cache_unit[Key_T, Value_T] {

	cacheUnit := new(cache_unit[Key_T, Value_T])

	cacheUnit.topic = topic
	cacheUnit.Logger = log.New(
		os.Stdout,
		fmt.Sprintf("[ CACHE TOPIC %s ]", cacheUnit.topic),
		0,
	)

	return cacheUnit
}

/*
Return the copy of the cached value
this method does not lock the cache value
*/
func (this *cache_unit[Key_T, Value_T]) Read(
	ctx context.Context, key Key_T,
) (value Value_T, exists bool, err error) {

	defer func() {

		if os.Getenv("CACHE_LOG") != "1" {

			return
		}

		go func() {
			switch {
			case err != nil:
				this.Println("reading key", key, "caused error:", err)
			case !exists:
				this.Println("reading inexisting key", key)
			default:
				this.Println("read", key)
			}
		}()
	}()

	cache, exists, e := this.getCache(ctx, key)

	if e != nil {

		err = e
		return
	}

	if !exists {

		return
	}

	lockAccquired, release, err := libTryLock.AccquireReadLock(ctx, &cache.RWMutex)

	defer func() {

		if lockAccquired {

			release()
		}
	}()

	if err != nil {

		return
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
func (this *cache_unit[Key_T, Value_T]) Set(ctx context.Context, key Key_T, value Value_T) (err error) {

	cache, exists, err := this.getCache(ctx, key)

	var oldVal Value_T

	defer func() {

		if os.Getenv("CACHE_LOG") != "1" {

			return
		}

		go func() {
			switch {
			case err != nil:
				this.Println("set key", key, "caused error:", err)
			case !exists:
				this.Println("key", key, "initiated by value", value)
			default:
				this.Println("key reassigned", key, "old value", oldVal, "new value", value)
			}
		}()
	}()

	if err != nil {

		return err
	}

	if !exists || cache.deactivated {

		cache = new(cache_value[Value_T])

		defer func() {

			if err != nil {

				this.Map.Swap(key, cache)
			}
		}()
	} else {

		oldVal = cache.value
	}

	err = this.set(ctx, cache, value)

	if err != nil {

		return err
	}

	// writeLog(fmt.Sprintf(`topic "%s" set key "%s"`, this.topic, key))
	fmt.Println("topic", this.topic, "set key", key)

	return nil
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

	var oldVal Value_T

	defer func() {

		if os.Getenv("CACHE_LOG") != "1" {

			return
		}

		go func() {
			switch {
			case err != nil:
				this.Println("set expiry key", key, "caused error:", err)
			case !exists:
				this.Println("key", key, "initiated by value", value, "until", moment)
			default:
				this.Println("key reassigned", key, "old value", oldVal, "new value", value, "until", moment)
			}
		}()
	}()

	if err != nil {

		return err
	}

	if !exists {

		currentCache = new(cache_value[Value_T])
		this.Map.Store(key, currentCache)

	} else {

		oldVal = currentCache.value
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

	if currentCache.expireTimer != nil {

		currentCache.expireTimer.Stop()
		currentCache.expireTimer = nil
	}

	cleanupTime := time.AfterFunc(duration, func() {

		deactivateCacheValue(currentCache)
		this.Delete(ctx, key)
		this.Println("key", key, "expired")
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

	defer func() {

		if os.Getenv("CACHE_LOG") != "1" {

			return
		}

		go func() {
			switch {
			case err != nil:
				this.Println("set key", key, "caused error:", err)
			case !exists:
				this.Println("deleting inexisting key", key)
			default:
				this.Println("key deleted", key, "old value", cache.value)
			}
		}()
	}()

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

	writeLockAcquired, releaseWrite, err := libTryLock.AcquireLock(ctx, cache)

	if err != nil {

		return
	}

	// readLockAcquired, releaseRead, err := libTryLock.AccquireReadLock(ctx, cache)

	// if err != nil {

	// 	return
	// }

	defer func() {

		if err != nil {

			releaseWrite()
		}
	}()

	if writeLockAcquired {

		updateCommand = func() (commit CommitFunction[Value_T], abort RevokeUpdateFunction) {

			commit = func(val Value_T) {

				cache.value = val
				//releaseRead()
				releaseWrite()
			}

			abort = func() {

				//releaseRead()
				releaseWrite()
			}

			return
		}

		return cache.value, updateCommand, nil
	}

	// if readLockAcquired {
	// 	releaseRead()
	// }

	// if writeLockAcquired {
	// 	releaseWrite()
	// }

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
