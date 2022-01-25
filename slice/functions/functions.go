package main

import "fmt"

func main() {
	var a = make([]int, 2, 2)
	// 改
	a[0] = 1
	a[1] = 2
	// 增
	a = append(a, 3)
	fmt.Println(a)
	// 删
	a = append(a[:1], a[2:]...)
	fmt.Println(a)
}
