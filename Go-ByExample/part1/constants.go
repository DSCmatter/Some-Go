package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func constants() {
	fmt.Println("This is" + " constants()")

	fmt.Println(s)

	const n = 500000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
