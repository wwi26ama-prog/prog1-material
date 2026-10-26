package main

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	settings_file        = "settings.txt"
	default_max_attempts = 3
	default_min_number   = 1
	default_max_number   = 100
)

func main() {
	max_attempts, min_number, max_number := readSettings()
	secret_number := rand.Intn(max_number-min_number+1) + min_number
	run(secret_number, max_attempts, min_number, max_number)
}

func readSettings() (int, int, int) {
	file, err := os.Open(settings_file)
	if err != nil {
		return default_max_attempts, default_min_number, default_max_number
	}
	defer file.Close()

	var max_attempts, min_number, max_number int
	_, err = fmt.Fscan(file, &max_attempts, &min_number, &max_number)
	if err != nil || max_attempts < 1 || min_number > max_number || max_number-min_number+1 <= 0 {
		return default_max_attempts, default_min_number, default_max_number
	}

	return max_attempts, min_number, max_number
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

func run(secret_number, max_attempts, min_number, max_number int) {
	for range max_attempts {
		var guess int
		fmt.Printf("Rate eine Zahl zwischen %d und %d: ", min_number, max_number)
		fmt.Scan(&guess)

		if checkGuess(guess, secret_number) == 0 {
			return
		}
	}
	fmt.Printf("Game Over! Die gesuchte Zahl war %d.\n", secret_number)
}
