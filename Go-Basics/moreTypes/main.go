// Main function to call all functions

package main

import (
	"fmt"
)

func main() {
	pointers()
	fmt.Println("This is struct", Vertex{1, 2})
	structField()
	PointerToStruct()
	fmt.Println("This is a struct literal:", v1, p, v2, v3)
	arrays()
	slice()
	sliceRefArray()
	sliceDefaults()
}
