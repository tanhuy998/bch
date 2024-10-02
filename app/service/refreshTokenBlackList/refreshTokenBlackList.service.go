package refreshTokenBlackListService

import (
	refreshTokenBlackListServicePort "app/port/refreshTokenBlackList"
	"context"
	"errors"
	"time"
)

var (
	ERR_REFRESH_TOKEN_ID_NOT_EXISTS = errors.New("RefreshTokenBlackListManipulatorService error: token id not exists")
	ERR_UPDATE_NOT_APPROVED         = errors.New("RefreshTokenBlackListManipulatorService error: update not approved")
)

type (
	IRefreshTokenBlackListManipulator = refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator
	IRefreshTokenBlackListPayload     = refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload

	// ReadFunction[Payload_T any]   func(payload Payload_T) error
	// UpdateFunction[Payload_T any] func(payload Payload_T) (newVal Payload_T, approve bool, err error)

	RefreshTokenBlackListManipulatorService struct {
		CacheClient refreshTokenBlackListServicePort.IRefreshTokenCacheClient // memoryCacheServicePort.IMemoryCacheClient[string, IRefreshTokenBlackListPayload]
	}
)

func (this *RefreshTokenBlackListManipulatorService) Has(tokenId string, ctx context.Context) (bool, error) {

	_, exists, err := this.CacheClient.Read(ctx, tokenId)

	if err != nil {

		return false, err
	}

	return exists, nil
}

func (this *RefreshTokenBlackListManipulatorService) Get(tokenID string, ctx context.Context) (IRefreshTokenBlackListPayload, bool, error) {

	return this.CacheClient.Read(ctx, tokenID)
}

func (this *RefreshTokenBlackListManipulatorService) Read(
	tokenID string,
	readFunc refreshTokenBlackListServicePort.ReadFunction,
	ctx context.Context,
) error {

	return this.CacheClient.Hold(ctx, tokenID, readFunc)
}

func (this *RefreshTokenBlackListManipulatorService) Update(
	tokenId string,
	updatefunc refreshTokenBlackListServicePort.UpdateFunction,
	ctx context.Context,
) error {

	this.CacheClient.Update(ctx, tokenId, updatefunc)

	return nil
}

func (this *RefreshTokenBlackListManipulatorService) Set(
	tokenId string,
	payload IRefreshTokenBlackListPayload,
	ctx context.Context,
) (bool, error) {

	err := this.CacheClient.Set(ctx, tokenId, payload)

	if err != nil {

		return false, err
	}

	return true, nil
}

func (this *RefreshTokenBlackListManipulatorService) SetWithExpire(
	tokenID string, payload IRefreshTokenBlackListPayload, expire time.Time, ctx context.Context,
) error {

	return this.CacheClient.SetWithExpire(ctx, tokenID, payload, expire)
}
