package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution.
// Goroutines are used to run concurrent threads...
// Or to be precise: A goroutine is a function running concurrently with other functions.

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutine1() {

	f("direct") // Suppose we have a function call f(s).
	// Here’s how we’d call that in the usual way, running it synchronously.

	go f("goroutine") // To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.

	go func(msg string) { //You can also start a goroutine for an anonymous function call.
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// Our two function calls are running asynchronously in separate goroutines now.
// Wait for them to finish (for a more robust approach, use a WaitGroup).

// When we run this program, we see the output of the blocking call first, then the output of the two goroutines.
// The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
