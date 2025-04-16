package main

import "fmt"

func main() {
	ch := make(chan int) // Creates a channel

	go func() {
		ch <- 42
	}()

	val := <-ch

	fmt.Println("Recieved:", val)
}
