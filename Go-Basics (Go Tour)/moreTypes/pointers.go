// Go has pointers. A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil.

package main

import (
	"fmt"
)

func pointers() {
	i, j := 43, 312
	fmt.Println("This is pointer function")

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // Now p points to 21
	fmt.Println(i, j)

	// Below is known as "dereferencing" or "indirecting".
	// Unlike C, Go has no pointer arithmetic.

	p = &j       // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

}
