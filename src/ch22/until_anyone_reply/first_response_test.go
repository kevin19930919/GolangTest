package untilanyonereply

import (
	"fmt"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("this is task %d", id)
}

func FirstResponse() string {
	numOfTasks := 10
	ch := make(chan string)
	for i := 0; i < numOfTasks; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log(FirstResponse())
}
