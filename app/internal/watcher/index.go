package watcher

import "sync"

var (
	wg sync.WaitGroup
)

func init() {

	wg = sync.WaitGroup{}
}

func Watch(fns ...func()) {

	for _, fn := range fns {

		wg.Add(1)

		go func() {

			fn()
			wg.Done()
		}()
	}
}

func Wait() {

	wg.Wait()
}
