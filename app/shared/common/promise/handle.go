package promise

import "context"

type (
	PromiseHandlerFunc[T any] func(context.Context, ResolveFunc[T], RejectFunc)
)

func Do[T any](ctx context.Context, fn PromiseHandlerFunc[T]) Promise[T] {

	if fn == nil {

		panic("")
	}

	ret := newPromiseObject[T](ctx)

	fn(ret.ctx, resolveFn(ret), rejectFn(ret))

	return ret
}
