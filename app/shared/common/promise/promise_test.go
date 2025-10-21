package promise

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPromise(t *testing.T) {

	sigFn := func(after time.Duration) <-chan time.Duration {
		c := make(chan time.Duration)
		go func() {

			defer close(c)
			time.Sleep(after)
			c <- after
		}()
		return c
	}

	resolveAfter := func(after time.Duration, resolve ResolveFunc[time.Duration], reject RejectFunc) Promise[time.Duration] {

		p, resolve, _ := New[time.Duration](context.TODO())

		<-sigFn(after)

		resolve(after)

		return p
	}

	p, resolve, reject := New[time.Duration](context.TODO())

	go resolveAfter(2*time.Hour, resolve, reject)
	go resolveAfter(5*time.Minute, resolve, reject)
	go resolveAfter(1*time.Second, resolve, reject)
	go resolveAfter(1*time.Hour, resolve, reject)
	go resolveAfter(1*time.Minute, resolve, reject)

	switch v, _, _ := p.Await(); {
	case v != 1*time.Second:
		t.Fail()
	default:
		fmt.Println(v)
	}
}

func TestResolvingPromise(t *testing.T) {

	p, resolve, _ := New[time.Duration](context.TODO())
	dur := 2 * time.Second

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {

		switch v, _, _ := p.Await(); {
		case v != dur:
			t.Fail()
		default:
			t.Logf("awaited duration %d", dur)
		}

		wg.Done()
	}()

	time.Sleep(dur)
	resolve(dur)
	t.Log("waiting")
	<-FutureOf(p)
	t.Log("promise settled")

	wg.Wait()
}

func TestRejectingPromise(t *testing.T) {

	p, _, reject := New[time.Duration](context.TODO())
	dur := 2 * time.Second

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {

		switch _, err, _ := p.Await(); {
		case err != nil:
			t.Logf("awaited duration %d", dur)
		}

		wg.Done()
	}()

	time.Sleep(dur)
	reject(fmt.Errorf("reject"))
	fmt.Println("waiting")
	<-FutureOf(p)
	t.Log("promise settled")

	wg.Wait()
}
