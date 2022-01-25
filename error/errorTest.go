package main

import "fmt"

func main() {
	defer func() {
		err := recover()
		if nil != err {
			fmt.Println(err)
		}
	}()
	a := 1
	b := 0
	fmt.Println(a / b)
}
