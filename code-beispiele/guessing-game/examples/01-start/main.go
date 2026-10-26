package main

import "fmt"

func main() {
	var guess int

	fmt.Print("Rate eine Zahl zwischen 1 und 100: ")
	fmt.Scan(&guess)

	if guess == 42 {
		fmt.Println("Richtig geraten!")
	} else {
		fmt.Println("Leider falsch geraten. Die gesuchte Zahl war 42.")
	}
}
