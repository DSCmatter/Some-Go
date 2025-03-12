// Learn how to control the flow of your code with conditionals, loops, switches and defers.
// Main func here...

package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	doubleit()
	while()
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(doThis(5))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	switch1()
	days()
	greeting()
	def()
	stackDef()
}
