package main

import (
	"sync"
)

func generator(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range numbers {
			out <- num
		}
		close(out)
	}()

	return out
}

func businessLogic(num int) int {
	return num * 2
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	var wg sync.WaitGroup

	ch1 := generator(data[0:3])
	ch2 := generator(data[3:])
	wg.Add(2)

	go func() {
		for num := range ch1 {
			println(businessLogic(num))
		}
		wg.Done()
	}()

	go func() {
		for num := range ch2 {
			println(businessLogic(num))
		}
		wg.Done()
	}()

	wg.Wait()
}
