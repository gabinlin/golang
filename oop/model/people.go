package model

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func (p People) Hello() {
	fmt.Println("Hello")
}

func NewPeople(name string, age int) People {
	return People{name, age}
}
