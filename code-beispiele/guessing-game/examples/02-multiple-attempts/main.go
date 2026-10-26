package main

import "fmt"

func main() {
	for range 3 {
		var guess int
		fmt.Print("Rate eine Zahl zwischen 1 und 100: ")
		fmt.Scan(&guess)

		if guess == 42 {
			fmt.Println("Richtig geraten!")
			return
		} else {
			fmt.Println("Leider falsch geraten.")
		}
	}
	fmt.Println("Game Over! Die gesuchte Zahl war 42.")
}
