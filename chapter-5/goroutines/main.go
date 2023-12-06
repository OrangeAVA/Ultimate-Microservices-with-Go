package main

import (
	"fmt"
	"time"
)

func main() {
	go printHello() // Hello

	go func() {
		fmt.Println("Hello from anonymous function")
	}() // Hello from anonymous function

	time.Sleep(1 * time.Second)
}

func printHello() {
	fmt.Println("Hello")
}
