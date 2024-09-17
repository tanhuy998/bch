package libTryLock

import (
	"context"
	"errors"
	"sync"
)

type (
	MutexPollingSignal  byte
	MutexPollingChannel chan MutexPollingSignal

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

const (
	StopPollingSignal       MutexPollingSignal = 0x0
	ContextDoneSignal       MutexPollingSignal = 0x1
	MutexLockAcquiredSignal MutexPollingSignal = 0x2
)

var (
	ERR_OUT_OF_CONTEXT = errors.New("polling: context done before lock acquired")
)

func tryRWMutexRLock(mutex ITryReadLockMutex) bool {

	return mutex.TryRLock()
}

func tryRWMutexLock(mutex ITryLockMutex) bool {

	return mutex.TryLock()
}

func waitLock[MutexMode_T any](
	ctx context.Context, mutex MutexMode_T, tryLockFunc TryRWLockFunction[MutexMode_T],
) {

	for ctx.Err() == nil {

		if tryLockFunc(mutex) {

			break
		}
	}
}

func pollRWMutex[MutexMode_T any](
	ctx context.Context, mutex MutexMode_T, tryFunc TryRWLockFunction[MutexMode_T], locker sync.Locker,
) (lockAcquired bool, err error) {

	if ctx == nil {

		ctx = context.TODO()
	}

	if any(mutex) == nil {

		panic("polling internal error: mutex is nil")
	}

	if tryFunc == nil {

		panic("polling internal error: try lock function is nil")
	}

	waitLock(ctx, mutex, tryFunc)

	if ctx.Err() != nil {

		locker.Unlock()
		err = ERR_OUT_OF_CONTEXT
		return
	}

	lockAcquired = true
	return
}

func AcquireLock(
	ctx context.Context, mutex ITryLockMutex,
) (lockAcquired bool, release func(), err error) {

	lockAcquired, err = pollRWMutex[ITryLockMutex](ctx, mutex, tryRWMutexLock, mutex)

	if err != nil {

		return
	}

	release = generateReleaseFunc(mutex)
	return
}

func AccquireReadLock(
	ctx context.Context, mutex ITryReadLockMutex,
) (lockAcquired bool, release func(), err error) {

	lockAcquired, err = pollRWMutex[ITryReadLockMutex](ctx, mutex, tryRWMutexRLock, mutex.RLocker())

	if err != nil {

		return
	}

	release = generateReleaseFunc(mutex.RLocker())
	return
}

func generateReleaseFunc(locker sync.Locker) func() {

	return func() {

		locker.Unlock()
	}
}
