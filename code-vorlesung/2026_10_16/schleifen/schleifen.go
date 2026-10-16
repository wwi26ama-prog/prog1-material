package main

import "fmt"

func main() {
	for range 3 {
		fmt.Println("A")
	}

	for i := 1; i < 4; i += 1 {
		fmt.Println(i)
	}

	for i := 0; i < 3; i += 1 {
		i = i + 1
		fmt.Println(i)
	}

	for j := range 3 {
		j = j + 1
		fmt.Println(j)
	}
}
