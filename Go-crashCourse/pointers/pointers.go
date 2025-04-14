package main

import "fmt"

func main() {
	a := 5
	b := &a // Assigns b to pointer of a, which is memeory address

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // *int -> (*) pointer of int

	// Use * to read value from address
	fmt.Println(*b) // -> GIves out value of &a which is 5

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
