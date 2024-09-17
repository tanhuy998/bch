package memoryCache

import (
	"context"
	"errors"
)

var (
	ERR_INVALID_CONTEXT    = errors.New("lock context initialization error: invalid parent context")
	ERR_LOCK_SAME_KEY      = errors.New("lock context error: locking the same key")
	ERR_DUBLICATE_lOCK_kEY = errors.New("lock context error: duplicate context key")
)

type (
	IKeyLockingContext[Key_T comparable] interface {
		context.Context
		LockKey() Key_T
		HasKey(key Key_T) bool
	}

	lock_context[Key_T, Value_T comparable] struct {
		context.Context
		keys     map[Key_T]struct{}
		value    Value_T
		lock_key Key_T
	}
)

func isKeyLockingContext[Key_T comparable](ctx any) bool {

	if _, isLockCtx := ctx.(IKeyLockingContext[Key_T]); isLockCtx {

		return true
	}

	return false
}

func newLockContext[Key_T, Value_T comparable](
	parentContext context.Context, key Key_T, value Value_T,
) (ctx *lock_context[Key_T, Value_T], err error) {

	var key_map map[Key_T]struct{}

	if !isKeyLockingContext[Key_T](parentContext) {

		key_map = make(map[Key_T]struct{})

	} else {

		v, _ := any(parentContext).(*lock_context[Key_T, Value_T])

		if key == v.lock_key {

			return nil, ERR_LOCKING_SAME_KEY
		}

		if v.HasKey(key) {

			return nil, ERR_DUBLICATE_lOCK_kEY
		}

		key_map = v.keys
	}

	key_map[key] = struct{}{}

	ret := &lock_context[Key_T, Value_T]{
		Context:  parentContext,
		lock_key: key,
		value:    value,
		keys:     key_map,
	}

	return ret, nil
}

func (this *lock_context[Key_T, Value_T]) LockKey() Key_T {

	return this.lock_key
}

func (this *lock_context[Key_T, Value_T]) LockValue() Value_T {

	return this.value
}

func (this *lock_context[Key_T, Value_T]) HasKey(key Key_T) bool {

	_, ok := this.keys[key]

	return ok
}
