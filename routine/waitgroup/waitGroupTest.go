package main

import (
	"fmt"
	"sync"
)

func main() {
	// 和java的countDown一个玩意
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func(n int) {
			fmt.Println(n)
			group.Done()
		}(i)
	}
	group.Wait()
}
