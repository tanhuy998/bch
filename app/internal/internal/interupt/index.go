package interupt

import (
	"os"
	"os/signal"
	"sync"
)

var (
	emited            bool
	ev_channel_list   []chan struct{} = make([]chan struct{}, 0)
	ev_channel_mutext sync.Mutex
)

func init() {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go poll(ch)
}

func poll(ch chan os.Signal) {

	defer close(ch)

	for range ch {

		emited = true
		emit_channels()
	}
}

func emit_channels() {

	ev_channel_mutext.Lock()
	defer ev_channel_mutext.Unlock()

	for _, c := range ev_channel_list {

		c <- struct{}{}
	}
}

func Emitted() bool {

	return emited
}

func C() <-chan struct{} {

	ev_channel_mutext.Lock()
	defer ev_channel_mutext.Unlock()

	ret := make(chan struct{})

	ev_channel_list = append(ev_channel_list, ret)

	return ret
}
