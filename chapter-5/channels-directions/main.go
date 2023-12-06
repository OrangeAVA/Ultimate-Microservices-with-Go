package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	send(ch)
	recv(ch)
}

func send(ch chan<- int) {
	ch <- 1
}

func recv(ch <-chan int) {
	fmt.Println(<-ch) // 1
}
