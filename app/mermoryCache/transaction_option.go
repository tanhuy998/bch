package memoryCache

import (
	libTryLock "app/lib/tryLock"
	"context"
	"errors"
)

var (
	ERR_READ_LOCK_NOT_ACQUIRED  = errors.New("read concern error: cannot acquire lock")
	ERR_WRITE_LOCK_NOT_ACQUIRED = errors.New("write concern error: cannot acquire lock")
)

type (
	transactionOption[Value_T any] interface {
		Activate(ctx context.Context, cache *cache_value[Value_T]) error
		Deactivate(cache *cache_value[Value_T])
		Err() error
	}

	readConcern[Value_T any] struct {
		//*sync.RWMutex
		suceed bool
	}

	writeConcern[Value_T any] struct {
		//*sync.RWMutex
		suceed bool
	}
)

func ReadConcernOption[Value_T any]() transactionOption[Value_T] {

	return new(readConcern[Value_T])
}

func WriteConcernOption[Value_T any]() transactionOption[Value_T] {

	return new(writeConcern[Value_T])
}

func (this *readConcern[Value_T]) Activate(ctx context.Context, cache *cache_value[Value_T]) error {

	readLockAcquired, _, err := libTryLock.AccquireReadLock(ctx, cache)

	if err != nil {

		return err
	}

	if !readLockAcquired {

		return ERR_READ_LOCK_NOT_ACQUIRED
	}

	this.suceed = true
	return nil
}

func (this *readConcern[Value_T]) Deactivate(cache *cache_value[Value_T]) {

	if !this.suceed {

		return
	}

	cache.RUnlock()
}

func (this *readConcern[Value_T]) Err() error {

	if !this.suceed {

		return ERR_READ_LOCK_NOT_ACQUIRED
	}

	return nil
}

func (this *writeConcern[Value_T]) Activate(ctx context.Context, cache *cache_value[Value_T]) error {

	readLockAcquired, _, err := libTryLock.AcquireLock(ctx, cache)

	if err != nil {

		return err
	}

	if !readLockAcquired {

		return ERR_WRITE_LOCK_NOT_ACQUIRED
	}

	this.suceed = true
	return nil
}

func (this *writeConcern[Value_T]) Deactivate(cache *cache_value[Value_T]) {

	if !this.suceed {

		return
	}

	cache.Unlock()
}

func (this *writeConcern[Value_T]) Err() error {

	if !this.suceed {

		return ERR_WRITE_LOCK_NOT_ACQUIRED
	}

	return nil
}
