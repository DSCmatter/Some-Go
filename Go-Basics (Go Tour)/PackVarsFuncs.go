// Packages, variables, and functions.
// Learn the basic components of any Go program.

package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool // All false by default
var i int                // 0 by default
var g, h = 1, 2

var a1 int = 11
var b2 float64 = float64(a1) // Type conversion
var c3 uint = uint(a1)

const Pi, Xi = 3.14, 1 // constant
const intNum64 = 412414144.64574512457553551

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	fmt.Println("hello")
	fmt.Println("My favorite number is", rand.Intn(100)) // random integer upto 100
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(16))
	fmt.Println(math.Pi)

	var l int
	j := l // Type inference

	fmt.Printf("j is of type %T\n", j)

	k := 3
	fmt.Println(add(32, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(i, c, python, java)
	fmt.Println(g, h, k)
	fmt.Println(a1, b2, c3)
	fmt.Println(intNum64) // Big int number

	fmt.Printf("This is a constant: Pi is %T, Xi is %T\n", Pi, Xi)
}
