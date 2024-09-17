package memoryCache

import (
	"context"
)

type (
	IHoldContext[Key_T, Value_T comparable] interface {
		IKeyLockingContext[Key_T]
		LockValue() Value_T
	}

	IUpdateContext[Key_T, Value_T comparable] interface {
		IHoldContext[Key_T, Value_T]
	}

	IModifyContext[Key_T, Value_T comparable] interface {
		//IUpdateContext[Key_T, Value_T]
		IUpdateContext[Key_T, Value_T]
		Commit(ctx context.Context, val Value_T)
	}

	ITransactionContext[Key_T, Value_T comparable] interface {
		context.Context
		IKeyLockingContext[Key_T]
		// Abort()
		// Commit(ctx context.Context, val Value_T)
	}

	transaction[Key_T, Value_T comparable] struct {
		context.Context
		*cache_value[Value_T]
		options  []transactionOption[Value_T]
		lock_key Key_T
		succeed  bool
	}
)

func startTransaction[Key_T, Value_T comparable](
	ctx context.Context, key Key_T, cache *cache_value[Value_T], options ...transactionOption[Value_T],
) (trans *transaction[Key_T, Value_T], err error) {

	if cache.deactivated {

		err = ERR_OBSOLETE_CACHE_VALUE
		return
	}

	// readLockAcquired, releaseRead, err := libTryLock.AccquireReadLock(ctx, cache)

	// if err != nil {

	// 	return
	// }

	// writeLockAcquired, releaseWrite, err := libTryLock.AcquireLock(ctx, cache)

	// if err != nil {

	// 	return
	// }

	ret := &transaction[Key_T, Value_T]{
		Context:     ctx,
		cache_value: cache,
		options:     options,
		lock_key:    key,
		//isAllLoksAcquired: readLockAcquired && writeLockAcquired,
	}

	// if !readLockAcquired || !writeLockAcquired {

	// 	if readLockAcquired {
	// 		releaseRead()
	// 	}

	// 	if writeLockAcquired {
	// 		releaseWrite()
	// 	}

	// 	err = ERR_OUT_OF_CONTEXT
	// 	return
	// }

	err = ret.activateOptions(ctx)

	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (this *transaction[Key_T, Value_T]) activateOptions(ctx context.Context) error {

	//succeed := true
	var err error = nil

	for _, v := range this.options {

		defer func() {

			if !this.succeed && v.Err() == nil {
				v.Deactivate(this.cache_value)
			}
		}()

		if v.Activate(ctx, this.cache_value) != nil {

			this.succeed = false
			err = v.Err()
			break
		}
	}

	return err
}

func (this *transaction[Key_T, Value_T]) deactivateOptions() {

	if !this.succeed {

		return
	}

	for _, v := range this.options {

		v.Deactivate(this.cache_value)
	}
}

func (this *transaction[Key_T, Value_T]) Abort() {

	this.deactivateTransaction()
}

func (this *transaction[Key_T, Value_T]) Commit(ctx context.Context, val Value_T) {

	if this.succeed {

		this.value = val
	}

	this.deactivateTransaction()
}

func (this *transaction[Key_T, Value_T]) deactivateTransaction() {

	if !this.succeed {

		return
	}

	this.deactivateOptions()
	// this.succeed = false
	// this.Unlock()
	// this.RUnlock()
}

func (this *transaction[Key_T, Value_T]) LockKey() Key_T {

	return this.lock_key
}

func (this *transaction[Key_T, Value_T]) LockValue() Value_T {

	return this.cache_value.value
}
