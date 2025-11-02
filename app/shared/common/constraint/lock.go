package constraint

import "sync"

type (
	NoCopy interface {
		sync.Locker
	}
)

type (
	IRMutex interface {
		RLock()
		RUnlock()
		TryRLock() bool
	}
)

type (
	IMutex interface {
		sync.Locker
		TryLock() bool
	}
)

type (
	IRWMutex interface {
		IMutex
		IRMutex
		RLocker() sync.Locker
	}
)
