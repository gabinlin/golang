package main

import "fmt"

func main() {
	var hello = "hello word"
	for i, i2 := range hello {
		fmt.Printf("%d:%c\n", i, i2)
	}
}
