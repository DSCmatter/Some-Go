// Slices
// An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
// flexible view into the elements of an array. In practice, slices are much more common than arrays.

// The type []T is a slice with elements of type T.

package main

import (
	"fmt"
)

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println("This is a slice:", s)
}

func sliceRefArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("This is slice like reference to arrays:", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// When slicing, you may omit the high or low bounds to use their defaults instead.
// The default is zero for the low bound and the length of the slice for the high bound.

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
