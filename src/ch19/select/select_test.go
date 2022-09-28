package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service(waitTiem int) string {
	fmt.Println("wait time: ", waitTiem)
	time.Sleep(time.Second * time.Duration(waitTiem))
	return "service done"
}

func AsyncServiceOne() chan string {
	retCh := make(chan string)
	go func() {
		ret := service(5)
		fmt.Println("service  one return")
		retCh <- ret
		fmt.Println("service one excited")
	}()
	return retCh
}

func AsyncServiceTwo() chan string {
	retCh := make(chan string)
	go func() {
		ret := service(10)
		fmt.Println("service two return")
		retCh <- ret
		fmt.Println("service two excited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case retCh1 := <-AsyncServiceOne():
		t.Log(retCh1)
	case retCh2 := <-AsyncServiceTwo():
		t.Log(retCh2)
	case <-time.After(time.Second * 7):
		t.Error("wait too long for service to finish...")
	}
}
