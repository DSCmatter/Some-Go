package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "apple"
	fruitArr[1] = "banana"
	// fruitArr[2] = "orange" -> Index out of bounds

	// Declare and assign at the same time
	VegetableArr := []string{"Radish", "Potatoes", "Ladyfinger"} // This is dynamic array since you can put any nos of elements without stating array index.

	fruitSlice := []string{"Apple", "Orange", "Grapes"}

	fmt.Println(fruitArr, fruitSlice)

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

	fmt.Println(VegetableArr)
	fmt.Printf("%T\n", fruitArr)
}
