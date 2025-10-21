package libTryLock

import (
	"context"
	"sync"
	"testing"
	"time"
)

// This test uses two goroutines to indicate that how LibtryLock performs.
// Main goroutine accquired the write lock aspect of the RWMutex
// until the read lock accquiring goroutine.
// This test success when the read lock accquiring goroutine deadline
// eslaped without read lock accquired.
func TestTryLockOnDeadlineContext(t *testing.T) {

	mu := sync.RWMutex{}

	writeLockAccquired, releaseWrite, err := AcquireLock(context.TODO(), &mu)

	switch {
	case err != nil:
		t.Fatal("error while accquiring write lock", err)
	case writeLockAccquired:
		t.Log("write lock accquired")
		defer releaseWrite()
	}

	t.Run("read lock accquire go routine", func(t *testing.T) {

		ctx, cancelCtx := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
		defer cancelCtx()

		t.Log("accquiring read lock for", 5*time.Second)
		startTime := time.Now()

		readLockAccquired, releaseRead, err := AccquireReadLock(ctx, &mu)

		switch {
		case err == nil:
			// err is used as signal to determine that the state of the accquisition
			t.Fatal("read lock accquiring goroutine must be return error")
		case readLockAccquired:
			t.Fatal("read lock accquired after", time.Until(startTime))
			defer releaseRead()
		default:
			t.Log("deadline for accquiring read lock runs out.")
		}
	})
}
