package main

import "runtime"

func forgottenSender(ch chan int) {
	data := 3
	// This is blocked as no one is receiving the data
	ch <- data
}

func main() {
	ch := make(chan int, 1) // avoid goroutine leak by using buffered channel
	go forgottenSender(ch)
	println(runtime.NumGoroutine()) // 2 instead of 1
}
