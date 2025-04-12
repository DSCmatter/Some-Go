package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y) // %d is placeholders for x & y
	} else {
		fmt.Printf("%d is greater than %d", x, y)
	}

	// else if
	color := "green"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not either blue or red")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not either blue or red")
	}
}
