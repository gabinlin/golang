package main

import "fmt"
import "awesomeProject/oop/model"

func main() {
	var p = model.People{"林嘉斌", 30}
	var p1 = model.People{Age: 29, Name: "林嘉斌"}
	var p2 = &model.People{Age: 31, Name: "吴敏"}

	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println(*p2)

	var teacher = model.Teacher{People: model.People{
		Name: "马士兵",
		Age:  45,
	}}
	fmt.Println(teacher)
}
