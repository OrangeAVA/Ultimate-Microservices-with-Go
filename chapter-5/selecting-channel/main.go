package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go func() {
		c1 <- 1
	}()
	go func() {
		c2 <- 2
	}()

	go func() {
		for {
			select {
			case <-c1:
				println("c1")
			case <-c2:
				println("c2")
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
