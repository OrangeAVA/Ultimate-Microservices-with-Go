package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func networkCall() int {
	time.Sleep(3 * time.Second)
	return 1
}

func forgottenSender(ch chan int) {
	data := networkCall()
	ch <- data
}

func handler() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	ch := make(chan int)
	go forgottenSender(ch)

	for {
		select {
		case data := <-ch:
			fmt.Printf("received data! %d\n", data)
			return nil
		case <-ctx.Done():
			return errors.New("timeout! Process canceled. Returning")
		}
	}
}

func main() {
	_ = handler()
}
