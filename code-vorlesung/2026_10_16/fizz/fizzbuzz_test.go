package main

import "fmt"

func ExampleGetOutput() {
	fmt.Println(GetOutput(1))
	fmt.Println(GetOutput(2))
	fmt.Println(GetOutput(3))
	fmt.Println(GetOutput(4))
	fmt.Println(GetOutput(5))
	fmt.Println(GetOutput(15))

	// Output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// FizzBuzz
}
