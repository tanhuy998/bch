package libPolling

import (
	libCommon "app/lib/common"
	"context"
	"fmt"
)

type (
	MutexPollingSignal  byte
	MutexPollingChannel chan MutexPollingSignal

	ITryLockMutex interface {
		TryLock() bool
	}

	ITryReadLockMutex interface {
		TryRLock() bool
	}

	TryRWLockFunction[MutexMode_t any] func(mutext MutexMode_t) bool

	IRWTryLockMutex interface {
		ITryLockMutex
		ITryReadLockMutex
	}
	/*
		WaitFunction blocks the current goroutine until
		either lock acquired or context Done()
		then return true or false respectedly
	*/
	WaitFunction func() bool
	/*
		ContextDoneCheckFuntion checks for the context have done yet
	*/
	ContextDoneCheckFunction func() bool
	/*
		StopPolling stop the polling loop in listening for
		signalChannel
	*/
	StopPollingFuncion func()

	mutex_polling_signal_bus struct {
		contextPollingChannel chan MutexPollingSignal
		contextDone           *bool
		stopPolling           bool
	}
)

const (
	StopPollingSignal       MutexPollingSignal = 0x0
	ContextDoneSignal       MutexPollingSignal = 0x1
	MutexLockAcquiredSignal MutexPollingSignal = 0x2
)

func tryRWMutexRLock(mutex ITryReadLockMutex) bool {

	return mutex.TryRLock()
}

func tryRWMutexLock(mutex ITryLockMutex) bool {

	return mutex.TryLock()
}

func pollContext(ctx context.Context, bus *mutex_polling_signal_bus) {

	// <-ctx.Done()
	// *contextDoneSignal = true

	defer fmt.Println("stop polling context")

	for !bus.stopPolling {

		select {
		case <-ctx.Done():
			//channel <- ContextDoneSignal
			*bus.contextDone = true
			bus.stopPolling = true
			//return
		case signal := <-bus.contextPollingChannel:
			/*
				listen for stop signal when context have not already Done().
				Some context would never Done() there for the goroutine will last
				for entire the program life cycle. when Lock acqured, the
			*/
			fmt.Println("signal", signal)
			if signal == StopPollingSignal {

				bus.stopPolling = true
			}
		}
	}
}

func pollRWMutexUnlock[MutexMode_T any](
	mutex MutexMode_T, bus *mutex_polling_signal_bus, tryLockFunc TryRWLockFunction[MutexMode_T],
) {

	defer fmt.Println("stop polling mutex")

	for !bus.stopPolling {
		switch {
		case *bus.contextDone:
			return
		case tryLockFunc(mutex):
			bus.contextPollingChannel <- MutexLockAcquiredSignal
			return
		}
	}
}

func pollRWMutex[MutexMode_T any](
	ctx context.Context, mutex MutexMode_T, tryFunc TryRWLockFunction[MutexMode_T],
) (wait WaitFunction, isContextDone ContextDoneCheckFunction, stopPolling StopPollingFuncion) {

	if ctx == nil {

		ctx = context.TODO()
	}

	if any(mutex) == nil {

		panic("polling internal error: mutex is nil")
	}

	if tryFunc == nil {

		panic("polling internal error: try lock function is nil")
	}

	pollingChannel := make(chan MutexPollingSignal)
	//contextDone := libCommon.PointerPrimitive(false)
	signalBus := &mutex_polling_signal_bus{
		contextPollingChannel: pollingChannel, // make(chan MutexPollingChannel),
		contextDone:           libCommon.PointerPrimitive(false),
		stopPolling:           false,
	}
	/*
		The polling process uses two goroutines, one for trying to acquire
		lock, one for listening to the context completion.
		there are communication between
	*/
	go pollContext(ctx, signalBus)
	go pollRWMutexUnlock(mutex, signalBus, tryFunc)

	return wrap(signalBus)
}

func wrap(
	bus *mutex_polling_signal_bus,
) (lockAccquired WaitFunction, contextDone ContextDoneCheckFunction, stopPolling StopPollingFuncion) {

	lockAccquired = func() bool {

		signal := <-bus.contextPollingChannel

		return signal == MutexLockAcquiredSignal
	}

	contextDone = func() bool {

		return *bus.contextDone
	}

	stopPolling = func() {

		stopMutexPolling(bus)
	}

	return
}

func AcquireLock(
	ctx context.Context, mutex ITryLockMutex,
) (wait WaitFunction, isContextDone ContextDoneCheckFunction, stopPolling StopPollingFuncion) {

	return pollRWMutex[ITryLockMutex](ctx, mutex, tryRWMutexLock)
}

func AccquireReadLock(
	ctx context.Context, mutex ITryReadLockMutex,
) (wait WaitFunction, isContextDone ContextDoneCheckFunction, stopPolling StopPollingFuncion) {

	return pollRWMutex[ITryReadLockMutex](ctx, mutex, tryRWMutexRLock)
}

func stopMutexPolling(bus *mutex_polling_signal_bus) {
	fmt.Println("stopping")
	//*doneSignal = true

	if !bus.stopPolling {

		bus.contextPollingChannel <- StopPollingSignal
	}

	fmt.Println("dispatch")
}
