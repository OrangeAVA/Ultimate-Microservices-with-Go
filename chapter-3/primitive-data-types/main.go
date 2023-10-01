package main

import "fmt"

func main() {
	var emptyMessage string
	fmt.Println(emptyMessage)

	var message string = "This is a message"
	fmt.Println(message)

	var inferredMessage = "This is an inferred message"
	fmt.Println(inferredMessage)

	alsoInferredMessage := "This is also an inferred message"
	fmt.Println(alsoInferredMessage)
}
