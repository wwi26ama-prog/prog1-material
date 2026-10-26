package main

import "fmt"

func Example_checkGuess_ok() {
	// Example usage of the CheckGuess function
	secret_number := 42
	guess := 42

	result := checkGuess(guess, secret_number)
	fmt.Println(result)

	// Output:
	// Richtig geraten!
	// 0
}

func Example_checkGuess_tooLow() {
	// Example usage of the CheckGuess function
	secret_number := 42
	guess := 30

	result := checkGuess(guess, secret_number)
	fmt.Println(result)

	// Output:
	// Die gesuchte Zahl ist größer.
	// -1
}

func Example_checkGuess_tooHigh() {
	// Example usage of the CheckGuess function
	secret_number := 42
	guess := 50

	result := checkGuess(guess, secret_number)
	fmt.Println(result)

	// Output:
	// Die gesuchte Zahl ist kleiner.
	// 1
}
