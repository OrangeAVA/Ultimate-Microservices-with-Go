package main

import "sync"

func merge(sourcesCh ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(sourcesCh))

	out := make(chan int)

	outputFunc := func(sourceCh <-chan int) {
		for num := range sourceCh {
			out <- num
		}
		wg.Done()
	}

	for _, sourceCh := range sourcesCh {
		go outputFunc(sourceCh)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

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

func main() {
	data := []int{1, 2, 3, 4, 5, 6}

	ch1 := generator(data[0:3])
	ch2 := generator(data[3:])

	ch := merge(ch1, ch2)

	for num := range ch {
		println(num)
	}
}
