package quizzes

// Eine Funktion, die ihren Parameter ändert, mit einer Zahl als Argument.

import "fmt"

// Double verdoppelt `n`.
func Double(n int) {
	n = n * 2
	fmt.Println(n)
}

// CallDouble ruft `Double` auf.
func CallDouble() {
	x := 3
	Double(x)
	fmt.Println(x)
}

// Frage:
// Was gibt der folgende Aufruf von `CallDouble` auf der Konsole aus?
// - `CallDouble()`
