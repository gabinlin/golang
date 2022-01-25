package main

import "fmt"

func main() {
	// finally 效果
	defer fmt.Println("你好")
	fmt.Println("Hello")

}
