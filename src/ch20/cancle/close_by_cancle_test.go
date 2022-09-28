package cancle

import (
	"fmt"
	"testing"
	"time"
)

func isCancled(cancleChan chan struct{}) bool {
	select {
	case <-cancleChan:
		return true
	default:
		return false
	}
}

func cancleOne(cancleChan chan struct{}) {
	cancleChan <- struct{}{}
}

func cancleTwo(cancleChan chan struct{}) {
	close(cancleChan)
}
func TestCancleOne(t *testing.T) {
	cancleChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancleChan chan struct{}) {
			for {
				if isCancled(cancleChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "channel is cancled")
		}(i, cancleChan)
	}
	cancleOne(cancleChan)
	time.Sleep(time.Second * 1)
}

func TestCancleTwo(t *testing.T) {
	cancleChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancleChan chan struct{}) {
			for {
				if isCancled(cancleChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "channel is cancled")
		}(i, cancleChan)
	}
	cancleTwo(cancleChan)
	time.Sleep(time.Second * 1)
}
