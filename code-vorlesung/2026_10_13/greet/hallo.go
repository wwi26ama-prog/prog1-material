package main

import "fmt"

func main() {
	fmt.Println(greet("Welt"))
}

func greet(name string) string {
	return fmt.Sprintf("Hallo %v", name)
}
