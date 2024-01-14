package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	go func() {
		select {
		case <-ctx.Done():
			println("Timeout")
			os.Exit(1)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
