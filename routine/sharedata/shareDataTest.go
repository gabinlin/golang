package main

import (
	"fmt"
	"sync"
)

func main() {
	var n = 0
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(t int) {
			n += t
			group.Done()
		}(i)
	}
	group.Wait()
	fmt.Println(n)
}
