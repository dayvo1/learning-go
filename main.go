package main

import "fmt"

func main() {
	fmt.Println("Starting my backend journey!")

	name := "David" // := is type inference (like auto in C++)
	age := 22

	fmt.Printf("I'm %s, %d years old\n", name, age)

	// Slices (like vectors in C++)
	skills := []string{"C++", "Java", "Learning Go"}
	for i, skill := range skills {
		fmt.Printf("%d. %s\n", i+1, skill)
	}
}
