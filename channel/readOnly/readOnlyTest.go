package main

import "fmt"

func main() {
	var ch = make(<-chan int, 3)
	// 只读，会报错
	//ch -> 1
	fmt.Println(<-ch)
}
