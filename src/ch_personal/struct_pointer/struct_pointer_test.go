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

// another test

type Grades struct {
	Math    int
	Science int
	Chinese int
	Total   func(...int) int
	// Average func(Math, Science, Chinese) int
}

// func TestAnonymousFuncInStruct(t *testing.T) {
// 	grades := &Grades{
// 		Math:    100,
// 		Science: 200,
// 		Chinese: 300,
// 		Total: func(...int) int {
// 			return Math + Science + Chinese
// 		},
// 	}
// 	t.Log(grades.Total)
// }

// func (g *Grades) Total() int {
// 	return g.Science + g.Chinese + g.Math
// }
