package memoryCacheServicePort

import (
	"app/internal/memoryCache"
	"context"
	"time"
)

type (
	IMemoryCacheClient[Key_T, Value_T comparable] interface {
		Read(ctx context.Context, key Key_T) (value Value_T, exists bool, err error)
		Hold(
			ctx context.Context, key Key_T, toDo func(ctx memoryCache.IHoldContext[Key_T, Value_T], value Value_T) error,
		) (err error)
		Set(ctx context.Context, key Key_T, value Value_T) error
		SetWithExpire(ctx context.Context, key Key_T, value Value_T, moment time.Time) error
		Delete(ctx context.Context, key Key_T) (deleted bool, err error)
		Update(
			ctx context.Context, key Key_T, toDo func(ctx memoryCache.IUpdateContext[Key_T, Value_T], val Value_T) (Value_T, error),
		) (err error)
	}
)
