package main

import (
	"time"
)

func main() {
	// has timeout
	resultCh := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		resultCh <- "Got before timeout"
	}()

	select {
	case res := <-resultCh:
		println(res)
	case <-time.After(2 * time.Second):
		println("Timeout")
	}

	// doesn't have timeout
	resultCh = make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		resultCh <- "Got before timeout"
	}()

	select {
	case res := <-resultCh:
		println(res)
	case <-time.After(2 * time.Second):
		println("Timeout")
	}
}
