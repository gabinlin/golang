package main

import "fmt"

func main() {
	// 方法1，从数组切出
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(arr[0:1]))
	// 方法2，make
	slice := make([]int, 2, 3)
	fmt.Println(&slice[0])
	slice[0] = 1
	fmt.Println(&slice[0])
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
