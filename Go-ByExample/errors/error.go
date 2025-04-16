package main

import (
	"errors"
	"fmt"
)

// Divide divides two integers and returns an error if b is zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero") // error message
	}
	return a / b, nil // successful result
}

func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
