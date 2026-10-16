package main

import "fmt"

func main() {
	for i := 1; i <= 24; i++ {
		fmt.Println(GetOutput(i))
	}
}

// GetOutput erwartet eine Zahl und liefert einen String:
// Entweder "Fizz", "Buzz" oder die Zahl selbst,
// entsprechend den Spielregeln.
func GetOutput(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	}
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprint(i)
}
