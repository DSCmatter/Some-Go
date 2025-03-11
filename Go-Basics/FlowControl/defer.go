// A defer statement defers the execution of a
// function until the surrounding function returns.

package main

import (
	"fmt"
)

func def() {
	defer fmt.Println("world")

	fmt.Println("hello") // execute this first
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.

func stackDef() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
