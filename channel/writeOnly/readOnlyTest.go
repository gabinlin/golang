package main

func main() {
	var ch = make(chan<- int, 3)
	ch <- 1
	// 只写，会报错
	//i := <-ch
}
