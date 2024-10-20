package cacheList

import (
	"app/internal/memoryCache"
	"context"
	"errors"
	"time"
)

var (
	ERR_REFRESH_TOKEN_ID_NOT_EXISTS = errors.New("CacheListManipulator[Key_T, Value_T]] error: token id not exists")
	ERR_UPDATE_NOT_APPROVED         = errors.New("CacheListManipulator[Key_T, Value_T]] error: update not approved")
)

type (
	CacheListManipulator[Key_T comparable, Value_T comparable] struct {
		CacheClient *memoryCache.CacheClient[Key_T, Value_T]
	}
)

func (this *CacheListManipulator[Key_T, Value_T]) Has(tokenId Key_T, ctx context.Context) (bool, error) {

	_, exists, err := this.CacheClient.Read(ctx, tokenId)

	if err != nil {

		return false, err
	}

	return exists, nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Get(tokenID Key_T, ctx context.Context) (Value_T, bool, error) {

	return this.CacheClient.Read(ctx, tokenID)
}

func (this *CacheListManipulator[Key_T, Value_T]) Read(
	tokenID Key_T,
	readFunc func(ctx memoryCache.IHoldContext[Key_T, Value_T], value Value_T) error,
	ctx context.Context,
) error {

	return this.CacheClient.Hold(ctx, tokenID, readFunc)
}

func (this *CacheListManipulator[Key_T, Value_T]) Update(
	tokenId Key_T,
	updatefunc func(ctx memoryCache.IUpdateContext[Key_T, Value_T], val Value_T) (Value_T, error),
	ctx context.Context,
) error {

	this.CacheClient.Update(ctx, tokenId, updatefunc)

	return nil
}

func (this *CacheListManipulator[Key_T, Value_T]) Set(
	tokenId Key_T,
	value Value_T,
	ctx context.Context,
) (bool, error) {

	err := this.CacheClient.Set(ctx, tokenId, value)

	if err != nil {

		return false, err
	}

	return true, nil
}

func (this *CacheListManipulator[Key_T, Value_T]) SetWithExpire(
	tokenID Key_T, value Value_T, expire time.Time, ctx context.Context,
) error {

	return this.CacheClient.SetWithExpire(ctx, tokenID, value, expire)
}

func (this *CacheListManipulator[Key_T, Value_T]) Delete(key Key_T, ctx context.Context) (bool, error) {

	return this.CacheClient.Delete(ctx, key)
}
