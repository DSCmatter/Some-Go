package main

import "fmt"

func arrays() {

	fmt.Println("This is" + " arrays()")

	var a [5]int
	fmt.Println("emp:", a) // Gives off an empty array

	a[4] = 100
	fmt.Println("set:", a)    // set a value at an index using the array[index] = value syntax
	fmt.Println("get:", a[4]) // Gets the value with array[index]

	fmt.Println("len:", len(a)) // len returns the length of an array

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b) // Use this syntax to declare and initialize an array in one line.

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dc1:", b) // You can also have the compiler count the number of elements for you with ...

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b) // If you specify the index with :, the elements in between will be zeroed.

	var twoD [2][3]int // Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// You can create and initialize multi-dimensional arrays at once too.

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
