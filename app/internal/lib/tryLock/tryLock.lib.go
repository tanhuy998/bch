package libTryLock

import (
	"context"
	"sync"
)

type (
	ITryLockMutex interface {
		TryLock() bool
		sync.Locker
		ILockReleaseMutex
	}

	ITryReadLockMutex interface {
		TryRLock() bool
		RLocker() sync.Locker
		IReadLockReleaseMutex
	}

	ILockReleaseMutex interface {
		Unlock()
	}

	IReadLockReleaseMutex interface {
		RUnlock()
	}

	TryRWLockFunction[MutexMode_t any] func(mutext MutexMode_t) bool
)

func tryRWMutexRLock(mutex ITryReadLockMutex) bool {

	return mutex.TryRLock()
}

func tryRWMutexLock(mutex ITryLockMutex) bool {

	return mutex.TryLock()
}

func waitLock[MutexMode_T any](
	ctx context.Context, mutex MutexMode_T, tryLockFunc TryRWLockFunction[MutexMode_T],
) (lockAccqured bool) {

	for ctx.Err() == nil {

		if tryLockFunc(mutex) {

			return true
		}
	}

	return false
}

func pollRWMutex[MutexMode_T any](
	ctx context.Context, mutex MutexMode_T, tryFunc TryRWLockFunction[MutexMode_T], locker sync.Locker,
) (err error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	if any(mutex) == nil {

		panic("polling internal error: mutex is nil")
	}

	if tryFunc == nil {

		panic("polling internal error: try lock function is nil")
	}

	lockAcquired := waitLock(ctx, mutex, tryFunc)

	defer func() {

		if lockAcquired && err != nil {
			locker.Unlock()
		}
	}()

	if ctx.Err() != nil {

		return ctx.Err()
	}

	return
}

func AcquireLock(
	ctx context.Context, mutex ITryLockMutex,
) (lockAcquired bool, release func(), err error) {

	err = pollRWMutex[ITryLockMutex](ctx, mutex, tryRWMutexLock, mutex)

	switch {
	case err != nil:
		return
	default:
		lockAcquired = true
		release = generateReleaseFunc(mutex)
		return
	}
}

func AccquireReadLock(
	ctx context.Context, mutex ITryReadLockMutex,
) (lockAcquired bool, release func(), err error) {

	err = pollRWMutex[ITryReadLockMutex](ctx, mutex, tryRWMutexRLock, mutex.RLocker())

	switch {
	case err != nil:
		return
	default:
		lockAcquired = true
		release = generateReleaseFunc(mutex.RLocker())
		return
	}
}

func generateReleaseFunc(locker sync.Locker) func() {

	return func() {

		locker.Unlock()
	}
}
