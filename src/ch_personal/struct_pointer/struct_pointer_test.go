package structpointer

import "testing"

type People struct {
	name string
	age  int
}

func NewPeople(name string, age int) *People {
	return &People{name: name, age: age}
}

func TestStruct(t *testing.T) {
	people := NewPeople("kevin", 28)
	t.Log("name: ", people.name, " age: ", people.age)
}
