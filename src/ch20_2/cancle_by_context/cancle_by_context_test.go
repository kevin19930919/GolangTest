package canclebycontext

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancleOne(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "channel is cancled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
