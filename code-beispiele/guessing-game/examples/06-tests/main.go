package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret_number := rand.Intn(100) + 1
	max_attempts := 3

	run(secret_number, max_attempts)
}

func checkGuess(guess, secret_number int) int {
	if guess == secret_number {
		fmt.Println("Richtig geraten!")
		return 0
	}
	if guess < secret_number {
		fmt.Println("Die gesuchte Zahl ist größer.")
		return -1
	} else {
		fmt.Println("Die gesuchte Zahl ist kleiner.")
		return 1
	}
}

func run(secret_number, max_attempts int) {
	for range max_attempts {
		var guess int
		fmt.Print("Rate eine Zahl zwischen 1 und 100: ")
		fmt.Scan(&guess)

		if checkGuess(guess, secret_number) == 0 {
			return
		}
	}
	fmt.Printf("Game Over! Die gesuchte Zahl war %d.\n", secret_number)
}
