package main

import "fmt"

func main() {
	fmt.Printf(`%d`, '2')
	greet("4" + "2")
	// greet(42) // Typfehler, 42 ist int, erwartet wird string
}

func greet(name string) {
	fmt.Printf("Hallo %v", name)
}
