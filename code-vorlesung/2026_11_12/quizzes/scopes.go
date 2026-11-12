package quizzes

// Eine Variable in verschachtelten Blöcken und Ausgaben auf der Konsole.

import "fmt"

// Scopes gibt mehrmals `x` aus.
func Scopes() {
	x := 1
	{
		x := 2
		{
			x = 3
			fmt.Println(x)
		}
		fmt.Println(x)
	}
	fmt.Println(x)
}

// Frage:
// Was gibt der folgende Aufruf von `Scopes` auf der Konsole aus?
// - `Scopes()`
