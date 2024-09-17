package libTryLock

// func KeepRead(ctx context.Context, wg *sync.WaitGroup, m *sync.RWMutex, label int, duration time.Duration) {
// 	fmt.Println(label)
// 	wg.Add(1)
// 	defer wg.Done()

// 	wait, isContextDone, stop, release := AcquireLock(ctx, m)

// 	defer stop()

// 	fmt.Println(isContextDone())

// 	acquired := wait()

// 	fmt.Println("acquired", acquired, "context Done", isContextDone())

// 	if acquired {
// 		fmt.Println("operation", label, "acquires write lock")
// 		defer release() // m.Unlock()
// 	}

// 	time.Sleep(duration)
// }

// func f() {

// 	m := &sync.RWMutex{}
// 	wg := &sync.WaitGroup{}

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
// 	cancel()

// 	KeepRead(ctx, wg, m, 1, time.Second*3)
// 	// go KeepRead(context.TODO(), wg, m, 2, time.Second)

// 	wg.Wait()

// 	go (func() {

// 		fmt.Println(2)
// 	})()

// 	fmt.Println("end")
// }

// func TestLock(t *testing.T) {

// 	f()
// }
