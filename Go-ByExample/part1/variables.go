package main

import (
	"fmt"
)

func vars() {
	fmt.Println("This is" + " vars()")
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e) // default value is 0

	f := "apple"
	fmt.Println(f)
}
