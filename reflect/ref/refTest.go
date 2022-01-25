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
	typeOf := reflect.TypeOf(people)
	valueOf := reflect.ValueOf(people)
	numField := typeOf.NumField()
	for i := 0; i < numField; i++ {
		field := typeOf.Field(i)
		fmt.Println(field.Name, valueOf.Field(i))
	}

}
