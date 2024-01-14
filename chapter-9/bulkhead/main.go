package main

import (
	"fmt"
	"time"

	"github.com/remeh/sizedwaitgroup"
)

func main() {
	swg := sizedwaitgroup.New(5)
	for i := 0; i < 100; i++ {
		swg.Add()
		go func(i int) {
			defer swg.Done()
			logic(i)
		}(i)
	}

	swg.Wait()
}

func logic(i int) {
	fmt.Println(i)
	time.Sleep(1 * time.Second)
}
