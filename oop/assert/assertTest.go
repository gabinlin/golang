package main

import (
	"awesomeProject/oop/interface"
	"awesomeProject/oop/model"
	"fmt"
)

func main() {
	var teacher _interface.Haha = model.Teacher{People: model.People{
		Name: "马士兵",
		Age:  45,
	}}
	m := teacher.(model.Teacher)
	fmt.Println(m)
}
