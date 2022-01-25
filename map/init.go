package main

import "fmt"

func main() {
	var map1 = make(map[string]int)
	map1["one"] = 1
	fmt.Println(map1)
	delete(map1, "one")
	fmt.Println(map1)
}
