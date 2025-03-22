package main

import "fmt"

// A struct is a collection of fields.

type Vertex struct {
	X, Y int
}

// Struct fields are accessed using a dot.

func structField() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println("This is struct field:", v.X)
}

// Pointers to structs
// Struct fields can be accessed through a struct pointer.

func PointerToStruct() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println("This is Pointer to Struct:", v)
}

// Struct Literals
// A struct literal denotes a newly allocated struct value by listing the values of its fields.

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
