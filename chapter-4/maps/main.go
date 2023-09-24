package main

import "fmt"

func main() {
	fruitsPrices := map[string]int{
		"orange": 10,
		"apple":  20,
		"banana": 15,
	}
	fmt.Println(fruitsPrices) // Output: map[apple:20 banana:15 orange:10]

	// Add a new key-value pair
	fruitsPrices["mango"] = 25
	fmt.Println(fruitsPrices) // Output: map[apple:20 banana:15 mango:25 orange:10]

	// Delete a key-value pair
	delete(fruitsPrices, "orange")
	fmt.Println(fruitsPrices) // Output: map[apple:20 banana:15 mango:25]

	// Get a value from a map
	fmt.Println(fruitsPrices["apple"]) // Output: 20

	// Get a value from a map that doesn't exist
	fmt.Println(fruitsPrices["grapes"]) // Output: 0 (the default value for int)

	// Check if a key exists in a map
	price, ok := fruitsPrices["grapes"]
	fmt.Println(price, ok) // Output: 0 false

	// Iterate over a map
	for fruit, price := range fruitsPrices {
		fmt.Printf("%s: %d\n", fruit, price)
	} // Output: apple: 20 banana: 15 mango: 25 (order may vary)

	fmt.Println(len(fruitsPrices)) // Output: 3

	// Create an empty map
	emptyMap := make(map[string]int)
	fmt.Println(emptyMap) // Output: map[]
}
