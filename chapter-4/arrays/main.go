package main

import "fmt"

func main() {
	// declare an array of 5 integers
	var a = [5]int{1, 2, 3, 4, 5}
	// declare an array with inferred length
	var b = [...]int{1, 2, 3}

	fmt.Println(a)
	fmt.Println(b)
}
