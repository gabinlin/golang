package main

import "fmt"

func main() {

	fmt.Println("哈哈")
	fmt.Println("哈哈")
	fmt.Println("哈哈")
	if 1 == 1 {
		goto label
	}
	fmt.Println("哈哈")
	fmt.Println("哈哈")
	fmt.Println("哈哈")
label:
	fmt.Println("哈哈")
	fmt.Println("哈哈")
}
