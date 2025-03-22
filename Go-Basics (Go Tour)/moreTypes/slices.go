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
	fmt.Println("This is sliceDefaults function...")
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceLengthCapacity() {
	fmt.Println("This is sliceLengthCapacity function...")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// The zero value of a slice is nil.
// A nil slice has a length and capacity of 0 and has no underlying array.

func nilSlice() {
	fmt.Println("This is a nilSlice: ")
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println(("nil!"))
	}
}

// Slices can be created with the built-in make function;
// this is how you create dynamically-sized arrays.
// The make function allocates a zeroed array and returns a slice that refers to that array:

func SliceMake() {
	fmt.Println("This is a SliceMake: ")
	a := make([]int, 5) // len(a)=5
	printSlice2("a", a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c) // len(b)=5, cap(b)=5

	d := c[2:5]
	printSlice2("d", d) // len(b)=4, cap(b)=4
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func AppendSlice() {
	fmt.Println("This is a AppendSlice: ")
	var x []int
	printSlice3(x)

	// append works on nil slices
	x = append(x, 0)
	printSlice3(x)

	// The slice grows as needed.
	x = append(x, 1)
	printSlice(x)

	// We can add more than one element at a time.
	x = append(x, 2, 3, 4)
	printSlice(x)
}

func printSlice3(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
