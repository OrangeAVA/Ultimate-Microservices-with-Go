package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			println(i)
		}(i)
	}

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Done")
}
