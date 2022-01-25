package model

import "fmt"

type Teacher struct {
	People
}

func (p Teacher) Hello() {
	fmt.Println("Hello")
}
