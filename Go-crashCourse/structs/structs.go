package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// Gender    string
	// age       int

	firstName, lastName, city, Gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.Gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	// Initialize person using struct
	// person1 := Person{firstName: "Simon", lastName: "Ozoumazi", city: "Montreal", Gender: "Male", age: 20}

	// Alternative
	person1 := Person{"Akari", "Ozoumazi", "Montreal", "Male", 20}

	person2 := Person{"Samantha", "Williams", "New York", "Female", 23}

	// Both gives same results

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName) // Get firstName value
	person1.age++                  // Increments age: 20 -> 21
	fmt.Println(person1)

	person1.hasBirthday()        // Increments again age as in hasBirthday()
	person1.getMarried("Thomas") // Akari gets married but doesn't change his lastName as he's a Male
	fmt.Println(person1.greet())

	person2.getMarried("Kaze")   // Samantha gets married and she changes her lastName as she's Female
	fmt.Println(person2.greet()) // Calls greet() to check lastName

}
