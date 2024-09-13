package refreshTokenBlackListService

import (
	memoryCacheServicePort "app/adapter/memoryCache"
	refreshTokenBlackListServicePort "app/adapter/refreshTokenBlackList"
	"context"
	"errors"
)

var (
	ERR_REFRESH_TOKEN_ID_NOT_EXISTS = errors.New("RefreshTokenBlackListManipulatorService error: token id not exists")
	ERR_UPDATE_NOT_APPROVED         = errors.New("RefreshTokenBlackListManipulatorService error: update not approved")
)

type (
	IRefreshTokenBlackListManipulator = refreshTokenBlackListServicePort.IRefreshTokenBlackListManipulator[string, IRefreshTokenBlackListPayload]
	IRefreshTokenBlackListPayload     = refreshTokenBlackListServicePort.IRefreshTokenBlackListPayload

	// ReadFunction[Payload_T any]   func(payload Payload_T) error
	// UpdateFunction[Payload_T any] func(payload Payload_T) (newVal Payload_T, approve bool, err error)

	RefreshTokenBlackListManipulatorService struct {
		CacheClient memoryCacheServicePort.IMemoryCacheClient[string, IRefreshTokenBlackListPayload]
	}
)

func (this *RefreshTokenBlackListManipulatorService) Has(tokenId string, ctx context.Context) (bool, error) {

	_, exists, err := this.CacheClient.ReadInstanctly(ctx, tokenId)

	if err != nil {

		return false, err
	}

	return exists, nil
}

func (this *RefreshTokenBlackListManipulatorService) Get(tokenID string, ctx context.Context) (IRefreshTokenBlackListPayload, bool, error) {

	return this.CacheClient.ReadInstanctly(ctx, tokenID)
}

func (this *RefreshTokenBlackListManipulatorService) Read(
	tokenID string,
	readFunc refreshTokenBlackListServicePort.ReadFunction[IRefreshTokenBlackListPayload],
	ctx context.Context,
) error {

	payload, exists, releaseLock, err := this.CacheClient.ReadAndHold(ctx, tokenID)

	if releaseLock != nil {

		defer releaseLock()
	}

	if err != nil {

		return err
	}

	if !exists {

		return ERR_REFRESH_TOKEN_ID_NOT_EXISTS
	}

	return readFunc(payload)
}

func (this *RefreshTokenBlackListManipulatorService) Update(
	tokenId string,
	updatefunc refreshTokenBlackListServicePort.UpdateFunction[IRefreshTokenBlackListPayload],
) error {

	payload, exists, command, err := this.CacheClient.Update(context.TODO(), tokenId)

	if err != nil {

		return err
	}

	if !exists {

		return ERR_REFRESH_TOKEN_ID_NOT_EXISTS
	}

	if updatefunc == nil {

		return nil
	}

	commit, revoke := command()

	newValue, approve, err := updatefunc(payload)

	if err != nil {

		return err
	}

	if !approve {

		revoke()
		return ERR_UPDATE_NOT_APPROVED
	}

	commit(newValue)
	return nil
}

func (this *RefreshTokenBlackListManipulatorService) Set(
	tokenId string,
	payload IRefreshTokenBlackListPayload,
) (bool, error) {

	err := this.CacheClient.Set(context.TODO(), tokenId, payload)

	if err != nil {

		return false, err
	}

	return true, nil
}
