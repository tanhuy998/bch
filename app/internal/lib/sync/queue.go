package sync

import "sync"

type (
	Queue[T any] struct {
		sync.Mutex
		queue []T
	}
)

func (this *Queue[T]) Init() {

	if this.IsEmpty() {

		this.queue = make([]T, 0)
	}
}

func (this *Queue[T]) Push(element T) {

	this.Lock()
	defer this.Unlock()

	this.queue = append(this.queue, element)
}

func (this *Queue[T]) IsEmpty() bool {

	return len(this.queue) == 0
}

func (this *Queue[T]) Pop() *T {

	if this.IsEmpty() {

		return nil
	}

	ret := this.queue[0]

	this.queue = append([]T{}, this.queue[1:len(this.queue)]...)

	return &ret
}

func (this *Queue[T]) Release() <-chan T {

	this.Lock()

	retChan := make(chan T, len(this.queue))

	for _, v := range this.queue {

		retChan <- v
	}

	go func() {

		for len(retChan) > 0 {

		}

		close(retChan)
		this.Unlock()
	}()

	return retChan
}
