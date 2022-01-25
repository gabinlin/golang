package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	var n = 0
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(t int) {
			mutex.Lock()
			n += t
			group.Done()
			mutex.Unlock()
		}(i)
	}
	group.Wait()
	fmt.Println(n)
}
