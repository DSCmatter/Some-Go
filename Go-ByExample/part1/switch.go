package main

import (
	"fmt"
	"time"
)

func switch1() {
	fmt.Println("This is" + " switch1()")

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	// In this example, the variable t will have the type corresponding to its clause.

	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Im a bool")
		case int:
			fmt.Println("Im an int")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	WhatAmI(true)
	WhatAmI(1)
	WhatAmI("hey")
}
