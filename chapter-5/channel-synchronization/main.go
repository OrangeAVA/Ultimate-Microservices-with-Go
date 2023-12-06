package main

import "time"

func main() {
	done := make(chan bool, 1)

	go func() {
		println("Start goroutine")
		time.Sleep(1 * time.Second) // simulate work
		println("End goroutine")
		done <- true
	}()

	println("Waiting goroutine")
	<-done

	println("All done")
}
