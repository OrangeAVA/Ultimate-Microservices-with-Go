package main

import "fmt"

func main() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}
