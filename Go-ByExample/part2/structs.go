package main

import "fmt"

// This person struct type has name and age fields.

type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name.

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func structs() {
	fmt.Println("This is a structs()")

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 21})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(s.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
