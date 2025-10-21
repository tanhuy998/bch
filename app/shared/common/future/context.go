package future

import (
	"context"
)

func SubstituteContextsWithCancel(contexts ...context.Context) (context.Context, context.CancelFunc) {

	ret, retCancel := context.WithCancel(context.TODO())

	superviseCtx, cancel := SubstituteContextsWithCancelCause(append(contexts, ret)...)

	go func() {

		//<-Race(channels...)
		<-superviseCtx.Done()

		cancel(__traceContextDone(append(contexts, ret)...))
	}()

	return ret, retCancel
}

func SubstituteContextsWithCancelCause(contexts ...context.Context) (context.Context, context.CancelCauseFunc) {

	ret, cancelCause := context.WithCancelCause(context.TODO())

	channels := make([]Future[struct{}], len(contexts))

	for i := range len(channels) {

		channels[i] = contexts[i].Done()
	}

	go func() {

		<-Race(channels...)
		cancelCause(__traceContextDone(contexts...))
	}()

	return ret, cancelCause
}

func SubstituteContexts(contexts ...context.Context) context.Context {

	ret, cancel := context.WithCancelCause(context.TODO())

	channels := make([]Future[struct{}], len(contexts))

	for i := range len(channels) {

		channels[i] = contexts[i].Done()
	}

	go func() {

		<-Race(channels...)
		cancel(__traceContextDone(contexts...))
	}()

	return ret
}

func __traceContextDone(contexts ...context.Context) (err error) {

	for _, ctx := range contexts {

		if ctx.Err() != nil {

			err = ctx.Err()
			return //ctx.Err()
		}
	}

	return nil
}
