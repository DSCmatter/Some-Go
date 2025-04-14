package main

import "fmt"

func main() {

	ids := []int{1, 2, 3, 4, 5}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Without using index -> _ represents i which is not here.
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// Range with maps
	emails := map[string]string{"Sike": "sike@gmail.com", "Chuck": "chuck@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
