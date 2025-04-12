package main

import (
	"Go-crashCourse/strutil"
	"fmt"
	"math"
)

func two() {
	fmt.Println("This is math()")

	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Sqrt(2))
	fmt.Println(math.Ceil(2.8))

	fmt.Println(strutil.Reverse("hello")) // This is imported from a package from /strutil
}
