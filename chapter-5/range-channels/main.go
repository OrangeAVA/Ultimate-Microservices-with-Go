package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch) // if you don't close the channel, deadlock will occur
	for i := range ch {
		fmt.Printf("%d ", i) // 0 1 2
	}
}
