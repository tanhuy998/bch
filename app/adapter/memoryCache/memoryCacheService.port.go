package memoryCacheServicePort

import (
	memoryCache "app/mermoryCache"
	"context"
)

type (
	IMemoryCacheClient[Key_T, Value_T comparable] interface {
		ReadInstanctly(ctx context.Context, key Key_T) (value Value_T, exists bool, err error)
		ReadAndHold(ctx context.Context, key Key_T) (value Value_T, exists bool, releaseLock memoryCache.ReadUnlockFunction, err error)
		Set(ctx context.Context, key Key_T, value Value_T) error
		Update(ctx context.Context, key Key_T) (value Value_T, keyExists bool, command memoryCache.UpdateCommandFunction[Value_T], err error)
		Delete(ctx context.Context, key Key_T) (deleted bool, err error)
	}
)
