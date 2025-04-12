package main

import "fmt"

var someName = "Pookie" // Global variable can be called from anywhere else

func one() {
	fmt.Println("This is one()")

	var name = "Brad"  // var name string = "Brad" - " " is already a string
	var age int32 = 37 // same with this as above.
	var isIdiot = true
	isIdiot = false // Now true -> false, vars can be redefined;

	const isStupid = true // so for unchanged variable use const

	// Shorthand
	Greetings := "Konnichwa" // can't be called outside functions
	size := 1.30             // float64 type

	Greetings1, size1 := "Konnichwa1", 1.301 // can also be called together

	fmt.Println(name, age, isIdiot, isStupid, someName, Greetings, size)
	fmt.Println(Greetings1, size1)
	fmt.Printf("%T\n", name)                                           // Print type of name
	fmt.Printf("%T %T %T %T %T\n", name, age, isIdiot, isStupid, size) // Print type of both of them

}
