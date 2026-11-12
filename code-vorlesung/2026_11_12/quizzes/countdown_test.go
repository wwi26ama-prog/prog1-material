package quizzes

func ExampleCountdown_five() {
	Countdown(5)

	// Output:
	// 5 odd
	// 4 even
	// 3 odd
	// 2 even
	// 1 odd
	// Go!
}

// Schleife mit fünf Durchläufen
// - (10) Die Schleife beginnt bei `start` und zählt mit `i--` herunter,
//   solange `i > 0` gilt. Die letzte ausgegebene Zahl ist also `1`.
// - (11) In jedem Durchlauf entscheidet `i%2 == 0`, ob `even` oder `odd`
//   hinter der Zahl steht.
// - (12, 14) `fmt.Println` trennt die Zahl und das Wort durch ein Leerzeichen.
// - (17) `Go!` steht nach der Schleife und wird deshalb genau einmal ausgegeben.

func ExampleCountdown_zero() {
	Countdown(0)

	// Output:
	// Go!
}

// Schleife ohne Durchlauf
// - (10) Bei `start` gleich `0` ist die Bedingung `i > 0` schon vor dem ersten
//   Durchlauf falsch, die Schleife wird also gar nicht durchlaufen.
// - (17) `Go!` steht nach der Schleife und wird trotzdem ausgegeben.
