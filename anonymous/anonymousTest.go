package main

import "fmt"

func main() {
	func() {
		fmt.Println("我是一个匿名函数")
	}()
}
