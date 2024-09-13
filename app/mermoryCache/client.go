package memoryCache

import (
	"context"
	"errors"
)

var (
	ERR_CACHE_TERMINATED       = errors.New("cache client error: cache terminated")
	ERR_CACHE_TOPIC_NOT_EXISTS = errors.New("cache client error: cache topic not exists")
)

type (
	CacheClient[Key_T, Value_T any] struct {
		topic string
	}
)

func NewClient[Key_T, Value_T any](topic string) (*CacheClient[Key_T, Value_T], error) {

	if cache_topics == nil {

		return nil, ERR_CACHE_TERMINATED
	}

	return &CacheClient[Key_T, Value_T]{topic: topic}, nil
}

/*
Return the copy of the cached value
this method does not lock the cache value
*/
func (this *CacheClient[Key_T, Value_T]) ReadInstanctly(ctx context.Context, key Key_T) (value Value_T, exists bool, err error) {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	return cacheUnit.ReadInstanctly(ctx, key)
}

/*
Return the copy of the cached value and lock the cache value
until the release lock func invoked,
if the cached value is not present
releaseLock func will be nil, ensure check of the second return value for
the existence of the cached value
*/
func (this *CacheClient[Key_T, Value_T]) ReadAndHold(
	ctx context.Context, key Key_T,
) (value Value_T, exists bool, releaseLock ReadUnlockFunction, err error) {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	return cacheUnit.ReadAndHold(ctx, key)
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

	return cacheUnit.Delete(ctx, key), nil
}

/*
Update a cached value, when the update method invoked, its locks until
the commitFunc commits the value to the cache room
*/
func (this *CacheClient[Key_T, Value_T]) Update(
	ctx context.Context, key Key_T,
) (value Value_T, keyExists bool, command UpdateCommandFunction[Value_T], err error) {

	cacheUnit, exists := GetTopic[Key_T, Value_T](this.topic)

	if !exists {

		err = ERR_CACHE_TOPIC_NOT_EXISTS
		return
	}

	return cacheUnit.Update(ctx, key)
}
