package main

import "fmt"

func main() {
	mySlice := []string{"nir", "david"}
	// mySlice := make([]type, length, capacity)
	fmt.Println(mySlice)

	slice := make([]int, 2, 5)
	fmt.Println(slice)      // [0 0]
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 5

	myArray := [5]int{1, 2, 3, 4, 5}
	slice = myArray[1:4] // from the 1st index to the 4th index (not included)
	fmt.Println(mySlice) // [2 3 4]

}
