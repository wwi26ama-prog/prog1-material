package quizzes

// Eine Funktion, die mit `append` an ihren Parameter, einen Slice, anhängt.

import "fmt"

// AddZero hängt eine `0` an `numbers` an.
func AddZero(numbers []int) {
	numbers = append(numbers, 0)
}

// ReallyAddZero hängt eine `0` an `numbers` an.
func ReallyAddZero(numbers []int) []int {
	numbers = append(numbers, 0)
	return numbers
}

// ReallyAddZero2 hängt eine `0` an `numbers` an.
func ReallyAddZero2(numbers *[]int) {
	*numbers = append(*numbers, 0)
}

// CallAddZero ruft `AddZero` auf.
func CallAddZero() {
	primes := []int{2, 3, 5}
	//AddZero(primes)
	//primes = ReallyAddZero(primes)
	ReallyAddZero2(&primes)
	fmt.Println(primes)
}

// Frage:
// Was gibt der folgende Aufruf von `CallAddZero` auf der Konsole aus?
// - `CallAddZero()`
