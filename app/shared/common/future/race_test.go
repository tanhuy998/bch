package future

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRace(t *testing.T) {

	wg := sync.WaitGroup{}

	sigFn := func(after time.Duration) <-chan time.Duration {
		c := make(chan time.Duration, 1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer close(c)
			time.Sleep(after)
			c <- after
		}()
		return c
	}

	dur := <-Race(
		sigFn(2*time.Second),
		sigFn(5*time.Second),
		sigFn(1*time.Second),
		sigFn(10*time.Second),
		sigFn(3*time.Second),
	)

	switch {
	case dur != 1*time.Second:
		t.Fail()
	default:
		t.Log(dur)
	}

	wg.Wait()
}

func TestRaceContext(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(6)
	listCtx := make([]context.Context, 6)

	mainCtx, mainCancel := context.WithCancelCause(context.TODO())

	listCtx[5] = mainCtx

	go func() {
		defer wg.Done()
		<-time.After(1 * time.Second)
		mainCancel(
			fmt.Errorf("main cancel context"),
		)
	}()

	for i := 0; i < 5; i++ {

		padding := 4

		err := fmt.Errorf("%v", time.Second*time.Duration(i+padding))

		ctx, cancel := context.WithDeadlineCause(context.TODO(), time.Now().Add(time.Second*time.Duration(i+padding)), err)

		go func() {
			<-ctx.Done()
			wg.Done()
		}()

		listCtx[i] = ctx
		defer cancel()
	}

	raceCtx := SubstituteContexts(listCtx...)

	<-raceCtx.Done()

	t.Log(raceCtx.Err())

	wg.Wait()
}

func TestRaceBaseOnMain(t *testing.T) {

	listCtx := make([]context.Context, 6)

	mainCtx, mainCancel := context.WithCancelCause(context.TODO())

	listCtx[5] = mainCtx

	go func() {

		<-time.After(2 * time.Second)
		mainCancel(
			fmt.Errorf("main cancel context"),
		)
	}()

	for i := 0; i < 5; i++ {

		padding := 3

		err := fmt.Errorf("%v", time.Second*time.Duration(i+padding))

		ctx, cancel := context.WithDeadlineCause(mainCtx, time.Now().Add(time.Second*time.Duration(i+padding)), err)

		listCtx[i] = ctx

		defer cancel()
	}

	raceCtx := SubstituteContexts(listCtx...)

	<-raceCtx.Done()

	t.Log(raceCtx.Err())
}
