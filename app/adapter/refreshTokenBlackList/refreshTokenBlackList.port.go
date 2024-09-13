package refreshTokenBlackListServicePort

import (
	"context"

	"github.com/google/uuid"
)

type (
	IRefreshTokenBlackListPayload interface {
		GetUserUUID() uuid.UUID
	}

	ReadFunction[Payload_T any]   func(payload Payload_T) error
	UpdateFunction[Payload_T any] func(payload Payload_T) (newVal Payload_T, approve bool, err error)

	IRefreshTokenBlackListManipulator[Key_T, Value_T any] interface {
		Has(tokenId Key_T, ctx context.Context) (bool, error)
		Get(tokenID Key_T) (IRefreshTokenBlackListPayload, bool, error)
		Read(tokenID Key_T, readFunc ReadFunction[IRefreshTokenBlackListPayload]) error
		Update(tokenId Key_T, updatefunc UpdateFunction[IRefreshTokenBlackListPayload]) error
		Set(tokenId Key_T, payload IRefreshTokenBlackListPayload) (bool, error)
	}
)
