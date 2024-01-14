package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if getEnvBool("feature-a", false) {
		fmt.Println("Feature A is enabled. Performing Feature A logic.")
		// Include Feature A logic here
	} else {
		fmt.Println("Feature A is not enabled.")
	}

	if getEnvBool("feature-b", false) {
		fmt.Println("Feature B is enabled. Performing Feature B logic.")
		// Include Feature B logic here
	} else {
		fmt.Println("Feature B is not enabled.")
	}

	// Rest of the application logic
}

func getEnvBool(flagName string, defaultValue bool) bool {
	value, exists := os.LookupEnv(flagName)
	if !exists {
		return defaultValue
	}

	result, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Printf("Error parsing environment variable %s: %v\n", flagName, err)
		return defaultValue
	}

	return result
}
