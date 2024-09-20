package refreshTokenBlackListServicePort

import (
	memoryCacheServicePort "app/adapter/memoryCache"
	memoryCache "app/mermoryCache"
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	REFRESH_TOKE_BLACK_LIST_TOPIC = "refresh-token-black-list"
)

type (
	IRefreshTokenCacheClient interface {
		memoryCacheServicePort.IMemoryCacheClient[string, IRefreshTokenBlackListPayload]
	}

	IRefreshTokenBlackListPayload interface {
		GetUserUUID() uuid.UUID
	}

	ReadFunction   func(ctx memoryCache.IHoldContext[string, IRefreshTokenBlackListPayload], payload IRefreshTokenBlackListPayload) error
	UpdateFunction func(ctx memoryCache.IUpdateContext[string, IRefreshTokenBlackListPayload], payload IRefreshTokenBlackListPayload) (newVal IRefreshTokenBlackListPayload, err error)

	IRefreshTokenBlackListManipulator interface {
		Has(tokenId string, ctx context.Context) (bool, error)
		Get(tokenID string, ctx context.Context) (IRefreshTokenBlackListPayload, bool, error)
		Read(tokenID string, readFunc ReadFunction, ctx context.Context) error
		Update(tokenId string, updatefunc UpdateFunction, ctx context.Context) error
		Set(tokenId string, payload IRefreshTokenBlackListPayload, ctx context.Context) (bool, error)
		SetWithExpire(tokenID string, payload IRefreshTokenBlackListPayload, exp time.Time, ctx context.Context) error
	}
)
