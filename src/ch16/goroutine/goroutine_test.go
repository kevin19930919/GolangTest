package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}

func TestCounterThreadSafe(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter)
}
