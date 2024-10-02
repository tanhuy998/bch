package memoryCache

import (
	"context"
	"errors"
	"time"
)

var (
	ERR_CACHE_TERMINATED       = errors.New("cache client error: cache terminated")
	ERR_CACHE_TOPIC_NOT_EXISTS = errors.New("cache client error: cache topic not exists")
	ERR_CACHE_KEY_NOT_EXISTS   = errors.New("cache client error: key not exists")
	ERR_TODO_FUNC_ABSENT       = errors.New("cache client error: toDo function is nil")
)

type (
	CacheClient[Key_T, Value_T comparable] struct {
		topic string
	}
)

func NewClient[Key_T, Value_T comparable](topic string) (*CacheClient[Key_T, Value_T], error) {

	if cache_topics == nil {

		return nil, ERR_CACHE_TERMINATED
	}

	return &CacheClient[Key_T, Value_T]{topic: topic}, nil
}

/*
Return the copy of the cached value
this method does not lock the cache value
*/
func (this *CacheClient[Key_T, Value_T]) Read(ctx context.Context, key Key_T) (value Value_T, exists bool, err error) {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	return cacheUnit.Read(ctx, key)
}

/*
Return the copy of the cached value and lock the cache value
until the release lock func invoked,
if the cached value is not present
releaseLock func will be nil, ensure check of the second return value for
the existence of the cached value
*/
func (this *CacheClient[Key_T, Value_T]) Hold(
	ctx context.Context, key Key_T, toDo func(ctx IHoldContext[Key_T, Value_T], value Value_T) error,
) (err error) {

	if toDo == nil {

		return ERR_TODO_FUNC_ABSENT
	}

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	if ctx.Err() != nil {

		return ERR_OUT_OF_CONTEXT
	}

	value, exists, release, err := cacheUnit.Hold(ctx, key)

	if err != nil {

		return err
	}

	if !exists {

		return ERR_CACHE_KEY_NOT_EXISTS
	}

	defer release()

	if ctx.Err() != nil {

		return ERR_OUT_OF_CONTEXT
	}

	isolateContext, err := newLockContext[Key_T, Value_T](ctx, key, value)

	if err != nil {

		return err
	}

	return toDo(isolateContext, value)
}

/*
Instantly Create a room for the value that corresponding to the key
if the cached key exists, the progress wiil lock for write operations
*/
func (this *CacheClient[Key_T, Value_T]) Set(ctx context.Context, key Key_T, value Value_T) error {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		return ERR_CACHE_TOPIC_NOT_EXISTS
	}

	return cacheUnit.Set(ctx, key, value)
}

/*
Instantly Delete the room of the cached value corresponding to the given key,
The delete progress locks on both read and write operations
*/
func (this *CacheClient[Key_T, Value_T]) Delete(ctx context.Context, key Key_T) (deleted bool, err error) {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	return cacheUnit.Delete(ctx, key)
}

/*
Update a cached value, when the update method invoked, its locks until
the commitFunc commits the value to the cache room
*/
func (this *CacheClient[Key_T, Value_T]) Update(
	ctx context.Context, key Key_T, toDo func(ctx IUpdateContext[Key_T, Value_T], val Value_T) (Value_T, error),
) (err error) {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	value, keyExists, command, err := cacheUnit.Modify(ctx, key)

	if err != nil {

		return
	}

	if !keyExists {

		err = ERR_CACHE_KEY_NOT_EXISTS
		return
	}

	commit, revoke := command()

	isolateContext, err := newLockContext(ctx, key, value)

	if err != nil {

		return err
	}

	updateValue, err := toDo(isolateContext, value)

	if err != nil {

		revoke()
		return err
	}

	commit(updateValue)
	return
}

func (this *CacheClient[Key_T, Value_T]) SetWithExpire(
	ctx context.Context, key Key_T, value Value_T, moment time.Time,
) error {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		return ERR_CACHE_TOPIC_NOT_EXISTS
	}

	return cacheUnit.SetWithExpire(ctx, key, value, moment)
}

// func (this *CacheClient[Key_T, Value_T]) Modify(
// 	ctx context.Context, key Key_T, toDo func(ctx IModifyContext[Key_T, Value_T], val Value_T) error,
// ) (err error) {

// 	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

// 	if !exists {

// 		err = ERR_CACHE_TOPIC_NOT_EXISTS
// 		return
// 	}

// 	cache, exists, err := cacheUnit.getCache(ctx, key)

// 	if err != nil {

// 		return err
// 	}

// 	if !exists {

// 		err = ERR_CACHE_KEY_NOT_EXISTS
// 		return
// 	}

// 	transaction, err := startTransaction(ctx, cache, ReadConcernOption[Value_T](), WriteConcernOption[Value_T]())

// 	if err != nil {

// 		return err
// 	}

// 	err = toDo(transaction, cache.value)

// 	if err != nil {

// 		transaction.Abort()
// 	}

// 	return err
// }
