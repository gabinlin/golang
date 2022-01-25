package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().AddDate(0, 1, 1))

	fmt.Printf("%T", time.Now())
}
