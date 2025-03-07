package main

import (
	"errors"
	"fmt"
)

// Type Enforcer funcs
func main() {
	var printValue string = "Hello World"
	PrintMe(printValue)

	var numerator int = 11
	var denominator int = 2
	// Capture the error returned from intDivision
	result, remainder, err := intDivision(numerator, denominator)
	if err != nil {
		// Handle the error
		fmt.Println(err.Error())
		return // exit the function if there's an error
	}
	fmt.Printf("Result: %d, Remainder: %d\n", result, remainder)
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

// Modified to return an error along with result and remainder
func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		// Fixed typo here: should be `errors.New`
		return 0, 0, errors.New("Cannot divide by zero")
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, nil
}
