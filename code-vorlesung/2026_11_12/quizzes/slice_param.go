package quizzes

// Eine Funktion, die ihren Parameter ändert, mit einem Slice als Argument.

import "fmt"

// DoubleAll verdoppelt die Zahlen in `numbers`.
func DoubleAll(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 2
	}
}

// CallDoubleAll ruft `DoubleAll` auf.
func CallDoubleAll() {
	primes := []int{2, 3, 5}
	DoubleAll(primes)
	fmt.Println(primes)
}

// Frage:
// Was gibt der folgende Aufruf von `CallDoubleAll` auf der Konsole aus?
// - `CallDoubleAll()`
