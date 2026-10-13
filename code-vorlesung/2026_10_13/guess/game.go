package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// alternativ: var secret int = ...
	secret := (rand.IntN(100) + 1)
	//secret := 42

	//for i := 0; i < 3; i = i + 1 {
	for range 3 {
		fmt.Print("Rate eine Zahl: ")

		var input int
		fmt.Scanln(&input)

		if check_input(secret, input) { // prüft, ob check_input true geliefert hat.
			return
		}
	}

	fmt.Println("Game Over")
}

func check_input(secret int, input int) bool {
	if input < secret {
		fmt.Println("Deine Zahl ist zu klein.")
		return false
	}

	if input > secret {
		fmt.Println("Deine Zahl ist zu groß.")
		return false
	}

	fmt.Println("Richtig geraten")
	return true
}
