package promise

import (
	"app/shared/common/future"
	"context"
	"sync"
)

func Race[T any](promises ...Promise[T]) Promise[T] {

	ret, resolve, reject := New[T](context.TODO())

	channels := make([]future.Future[interface{}], len(promises))

	for i := range len(promises) {

		channels[i] = future.Future[interface{}](FutureOf(promises[i]))
	}

	go func() {

		switch v := (<-future.Race(channels...)).(type) {
		case T:
			resolve(v)
		case error:
			reject(v)
		}
	}()

	return ret
}

func All[T any](promises ...Promise[T]) Promise[[]T] {

	ret, resolve, reject := withResolver[[]T](context.TODO())
	stop := make(chan struct{}, 1)

	fulfillments := make([]T, len(promises))
	fulfillCount := 0

	for i, p := range promises {

		go func() {

			select {
			case v := <-FutureOf(p):
				switch actualVal := v.(type) {
				case T:
					fulfillments[i] = actualVal
					fulfillCount++
				case error:
					reject(actualVal)
					stop <- struct{}{}
					close(stop)
				}
				return
			case <-stop:
				return
			}
		}()
	}

	go func() {

		for _, stopped := <-stop; !stopped || fulfillCount < len(promises) && ret.StillHolds(); {

		}

		switch {
		case fulfillCount == len(promises):
			resolve(fulfillments)
		}
	}()

	return ret
}

func AllSettled[T any](promises ...Promise[T]) Promise[[]interface{}] {

	ret, resolve, _ := New[[]interface{}](context.TODO())

	settlements := make([]interface{}, len(promises))
	wg := sync.WaitGroup{}

	for i, p := range promises {

		wg.Add(1)

		go func() {

			defer wg.Done()

			settlements[i] = <-FutureOf(p)
		}()
	}

	go func() {

		wg.Wait()
		resolve(settlements)
	}()

	return ret
}
