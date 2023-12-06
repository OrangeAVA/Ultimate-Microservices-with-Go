package main

func main() {
	ch := make(chan int)
	close(ch)
	ch <- 1 // panic: send on closed channel
}
