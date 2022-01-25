package main

import (
	"awesomeProject/oop/model"
	"fmt"
	"reflect"
)

func main() {
	people := model.People{
		Name: "哈哈",
		Age:  0,
	}
	of := reflect.ValueOf(people)
	i := of.Interface()
	m := i.(model.People)
	fmt.Println(m)

	var a = 1
	valueOf := reflect.ValueOf(a)
	i3 := valueOf.Interface()
	i2 := i3.(int)
	fmt.Println(i2)
	fmt.Printf("%T", valueOf.Type())
}
