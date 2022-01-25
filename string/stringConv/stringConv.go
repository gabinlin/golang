package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 1
	var msg = fmt.Sprintf("%d", i)
	fmt.Println(msg)

	var j, _ = strconv.ParseInt(msg, 10, 0)
	fmt.Println(j)

}
