package main

import "fmt"

func exchange(a *int, b *int) {
	var c = *a
	*a = *b
	*b = c
}

func main() {
	var a = 1
	var b = 2
	exchange(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
