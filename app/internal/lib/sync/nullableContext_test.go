package sync_test

import (
	libCommon "app/internal/lib/common"
	"app/internal/lib/sync"
	"context"
	"testing"
	"time"
)

func TestNullableContext(t *testing.T) {

	ctx := new(sync.NullableContext[context.Context])

	signal := 0

	go func() {

		ctx.Deadline()

		signal = 1
	}()

	go func() {

		<-time.After(time.Second * 5)

		ctx.SetBaseContext(libCommon.PointerPrimitive(context.TODO()))
	}()

	<-time.After(time.Second * 10)

	if signal == 0 {

		t.Error("failed")
	}
}
