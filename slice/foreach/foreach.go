package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for _, val := range arr {
		fmt.Println(val)
	}
}
