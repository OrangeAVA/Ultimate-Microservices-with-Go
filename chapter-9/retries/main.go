package main

import (
	"fmt"
	"time"
)

func main() {
	err := Retry(3, time.Second, func() error {
		fmt.Println("trying...")
		return fmt.Errorf("some error")
	})
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	currentAttempt := 1
	for {
		err = f()
		if err == nil {
			return nil
		}
		if currentAttempt >= attempts {
			return fmt.Errorf("after %d attempts, last error: %w", attempts, err)
		}
		time.Sleep(sleep)
		sleep *= 2
		currentAttempt++
	}
}
