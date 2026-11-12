package quizzes

// Aufwärmübung: eine Schleife mit einer Verzweigung und Ausgaben auf der Konsole.

import "fmt"

// Countdown erwartet eine ganze Zahl `start` und gibt einige Zeilen aus.
func Countdown(start int) {
	for i := start; i > 0; i-- {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	fmt.Println("Go!")
}

// Frage:
// Was geben die folgenden beiden Aufrufe von `Countdown` auf der Konsole aus?
// - `Countdown(5)`
// - `Countdown(0)`
