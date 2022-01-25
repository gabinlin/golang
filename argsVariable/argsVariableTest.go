package main

import (
	"awesomeProject/math"
	"fmt"
)

func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println(math.Add(1, 2, 4))
}
