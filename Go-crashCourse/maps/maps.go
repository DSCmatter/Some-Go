package main

import "fmt"

func main() {
	// Define map

	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"]) // Call value of Bob
	fmt.Printf("%T\n", emails)

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add keys

	emails1 := map[string]string{"Sike": "sike@gmail.com", "Chuck": "chuck@gmail.com"}
	emails1["Jerry"] = "jerry@gmail.com"
	fmt.Println(emails1)

}
