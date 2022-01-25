package main

import (
	"fmt"
	"time"
)

func main() {
	var ch1 = make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		for {
			select {
			case t := <-ch1:
				fmt.Println(t)
			}
		}
	}()
	time.Sleep(time.Second * 2)
}
