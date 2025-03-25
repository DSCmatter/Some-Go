package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func MultipleReturnVals() {
	fmt.Println("This is" + " MultipleReturnVals()")

	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := vals()
	fmt.Println("c:", c)
}
