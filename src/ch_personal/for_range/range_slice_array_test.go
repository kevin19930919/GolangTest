package for_range

import (
	"testing"
)

func TestRangeArray(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	for index, value := range a {
		if index == 0 {
			a[1] = 10
		}
		t.Log("index: ", index, "value: ", value, "value pointer: ", &value)
	}
}

func TestRangeSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	for index, value := range a {
		if index == 0 {
			a[1] = 10
		}
		t.Log("index: ", index, "value: ", value, "value pointer: ", &value)
	}
}

func TestRangeSliceAppend(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	var b []int
	for index, value := range a {
		b = append(b, value)
		t.Log("index: ", index, "value: ", value, "value pointer: ", &value, "b: ", b)
	}
	t.Log("final result: ", b)
}
