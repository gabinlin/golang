package main

func main() {
	var ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
}
