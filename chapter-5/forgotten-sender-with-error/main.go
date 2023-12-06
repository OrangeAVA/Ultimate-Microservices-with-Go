package main

import (
	"errors"
	"fmt"
)

func networkCall() int {
	return 1
}

func anotherAction() error {
	return errors.New("data is invalid! Returning")
}

func forgottenSender(ch chan int) {
	data := networkCall()

	ch <- data
}

func handler() error {
	ch := make(chan int)
	go forgottenSender(ch)

	err := anotherAction()
	if err != nil {
		return err
	}

	data := <-ch
	fmt.Println(data)
	return nil
}

func main() {
	_ = handler()
}
