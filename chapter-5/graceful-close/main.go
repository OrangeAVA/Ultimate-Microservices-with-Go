package main

import "sync"

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Sender(ch, wg, i)
	}

	go func() {
		wg.Wait()
		close(ch) // graceful close
	}()

	// receiver
	for i := range ch {
		println(i)
	}
}

func Sender(ch chan<- int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	ch <- i
}
