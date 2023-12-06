package main

import "fmt"

func leak(ch chan int) {
	data := <-ch
	fmt.Println(data)
}

func main() {
	ch := make(chan int)

	go leak(ch)
}
