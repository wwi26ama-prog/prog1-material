package main

import "fmt"

func main() {
	secret_number := 42
	max_attempts := 3

	run(secret_number, max_attempts)
}

func run(secret_number, max_attempts int) {
	for range max_attempts {
		var guess int
		fmt.Print("Rate eine Zahl zwischen 1 und 100: ")
		fmt.Scan(&guess)

		if guess == secret_number {
			fmt.Println("Richtig geraten!")
			return
		} else {
			fmt.Println("Leider falsch geraten.")
		}
	}
	fmt.Printf("Game Over! Die gesuchte Zahl war %d.\n", secret_number)
}
