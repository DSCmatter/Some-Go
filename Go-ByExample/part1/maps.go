package main

import (
	"fmt"
	"maps"
)

func SomeMap() {
	fmt.Println("This is" + " maps()")

	m := make(map[string]int) // make(map[key-type]val-type).

	m["k1"] = 7 // Set key/value pairs using typical name[key] = val syntax.
	m["k2"] = 13

	fmt.Println("map:", m) // Printing a map with e.g. fmt.Println will show all of its key/value pairs.

	v1 := m["k1"] // Get a value for a key with name[key].
	fmt.Println("v1:", v1)

	v3 := m["k3"] // If a value doesnt exist it returns 0
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m) // To remove all key/value pairs from a map
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n and n2 are equal")
	} else {
		fmt.Println("n and n2 are not equal")
	}

}
