package promise

import (
	"app/shared/common/future"
	"fmt"
)

type (
	PromiseFuture future.Future[interface{}]
)

func BelieveIn[T any](p Promise[T]) PromiseFuture {

	return FutureOf(p)
}

func FutureOf[T any](p Promise[T]) PromiseFuture {

	switch _p := p.(type) {
	case *promise[T]:
		return _p._future()
	default:
		ret := make(chan interface{}, 1)

		ret <- fmt.Errorf("could not see future of the given promise")
		close(ret)
		return ret
	}
}
