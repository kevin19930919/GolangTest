package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Second * 1)
	return "service done"
}

func otherTask() {
	fmt.Println("working on something .......")
	time.Sleep(time.Second * 3)
	fmt.Println("finish task...")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("service return")
		retCh <- ret
		fmt.Println("service excited")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 3)
}
