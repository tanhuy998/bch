package libCommon

import "context"

func pollContext(ctx context.Context, signal *int) (stopPolling chan interface{}) {

	stopPolling = make(chan interface{}, 1)
	go pollContextFunc(ctx, stopPolling)
	return
}

func pollContextFunc(ctx context.Context, stopPolling chan interface{}) {

	for {
		select {
		case <-ctx.Done():
			stopPolling <- true
		case <-stopPolling:
			return
		}
	}
}
