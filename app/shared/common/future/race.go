package future

type (
	IFuture[T any] interface {
		Await() T
	}
)

type (
	Future[T any] <-chan T
)

func (f Future[T]) Value(key any) any {

	return nil
}

func Race[T any](channels ...Future[T]) Future[T] {

	switch len(channels) {
	case 0:
		ret := make(chan T)
		close(ret)
		return ret
	case 1:
		return Future[T](channels[0])
	default:
		ret := make(chan T) // avoid deadlock on read

		go __poll(ret, channels...)

		return Future[T](ret)
	}
}

func __poll[T any](base chan T, channels ...Future[T]) {

	defer close(base)

	var v T

	select {
	case v = <-channels[0]:
	case v = <-channels[1]:
	case v = <-Race(append(channels[2:], base)...):
	}

	base <- v
}
