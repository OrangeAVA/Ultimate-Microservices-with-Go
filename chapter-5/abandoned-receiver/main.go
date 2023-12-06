package main

import (
	"runtime"
	"time"
)

func abandonedReceiver(ch chan string) {
	for data := range ch {
		println(data)
	}

	println("Worker is done")
}

func sender(ch chan string) {
	for _, data := range []string{"one", "two", "three"} {
		ch <- data
	}
	// close(ch) // solve the goroutine leak
}

func handler() {
	ch := make(chan string, 3)
	sender(ch)

	go abandonedReceiver(ch)

}

func main() {
	handler()
	time.Sleep(1 * time.Second)

	println(runtime.NumGoroutine())
}
