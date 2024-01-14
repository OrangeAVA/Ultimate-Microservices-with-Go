package main

import (
	"fmt"
)

func main() {
	result, err := performOperation()
	if err != nil {
		// Fallback operation in case of an error
		result = fallbackOperation()
	}

	fmt.Println("Result:", result)
}

func performOperation() (string, error) {
	return "", fmt.Errorf("Operation failed")
}

func fallbackOperation() string {
	return "Fallback result"
}
