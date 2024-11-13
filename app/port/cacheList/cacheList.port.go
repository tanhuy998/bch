package cacheListServicePort

import (
	"app/internal/memoryCache"
	"context"
	"time"
)

type (
	ICacheList[Key_T, Value_T comparable] interface {
		Has(tokenId Key_T, ctx context.Context) (bool, error)
		Get(tokenID Key_T, ctx context.Context) (Value_T, bool, error)
		Read(
			tokenID Key_T,
			readFunc func(ctx memoryCache.IHoldContext[Key_T, Value_T], value Value_T) error,
			ctx context.Context,
		) error
		Update(
			tokenId Key_T,
			updatefunc func(ctx memoryCache.IUpdateContext[Key_T, Value_T], val Value_T) (Value_T, error),
			ctx context.Context,
		) error
		Set(
			tokenId Key_T,
			value Value_T,
			ctx context.Context,
		) (bool, error)
		SetWithExpire(
			tokenID Key_T, value Value_T, expire time.Time, ctx context.Context,
		) error
		Delete(key Key_T, ctx context.Context) (bool, error)
	}
)
