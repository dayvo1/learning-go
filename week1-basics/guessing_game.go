package main

import (
	"fmt"
	"math/rand"
)

func guessingGame() {
	target := rand.Intn(100) + 1

	fmt.Println("Guess a number between 1 and 100")

	attempts := 0

	for {
		var guess int
		fmt.Print("Your guess: ")
		fmt.Scan(&guess)
		if guess == target {
			// print you win
			fmt.Printf("You won in %v attempts", attempts)
			break
		} else if guess < target {
			attempts++
			fmt.Println("Too low, try again")
		} else {
			attempts++
			fmt.Println("Too high, try again")
		}
	}

}
