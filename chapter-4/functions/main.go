package main

import "fmt"

func main() {
	fmt.Println(myFunction(5, "nir"))
	fmt.Println(multipleReturns("nir", 5))
	fmt.Println(namedReturnValues("nir", 5))

	functionWithDefer()
	functionWithMultipleDefers()
}

func myFunction(amount int, prefix string) string {
	return fmt.Sprintf("%s: %d", prefix, amount)
} // Input: 5, "nir" | Output: "nir: 5"

func multipleReturns(prefix string, amount int) (string, int) {
	amount++
	return fmt.Sprintf("%s: %d", prefix, amount), amount
} // Input: "nir", 5 | Output: "nir: 6", 6

func namedReturnValues(prefix string, amount int) (result string, newAmount int) {
	amount++
	result = fmt.Sprintf("%s: %d", prefix, amount)
	newAmount = amount
	return
} // Input: "nir", 5 | Output: "nir: 6", 6

func functionWithDefer() {
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
} // Output: This will be printed first
//         This will be printed last

func functionWithMultipleDefers() {
	defer func() {
		fmt.Println("This will be printed last")
	}()
	defer fmt.Println("This will be printed second")
	fmt.Println("This will be printed first")
}

// Output: This will be printed first
//         This will be printed second
//         This will be printed last
