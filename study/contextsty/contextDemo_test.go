package contextsty

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextWithCancelDemo(t *testing.T) {
	ContextWithCancelDemo()
}

func TestContextWithTimeOutDemo(t *testing.T) {
	ContextWithTimeOutDemo()
}

func TestWaitGroupDemo(t *testing.T) {
	WaitGroupDemo0()
	WaitGroupDemo1()
}

func TestOther(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	go func(cancel context.CancelFunc) {
		time.Sleep(1 * time.Second)
		cancel()
	}(cancelFunc)
	select {
	case <-ctx.Done():
		fmt.Sprint("done")
	}
}
