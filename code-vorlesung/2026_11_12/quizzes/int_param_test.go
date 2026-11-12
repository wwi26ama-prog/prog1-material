package quizzes

func ExampleCallDouble() {
	CallDouble()

	// Output:
	// 6
	// 3
}

// - (15) `x` ist eine Variable mit dem Wert `3`.
// - (9, 16) Beim Aufruf bekommt der Parameter `n` eine Kopie des Werts
//   von `x`. `n` und `x` sind zwei verschiedene Variablen.
// - (10) `Double` ändert nur seine Kopie `n` von `3` auf `6`.
// - (17) `x` ist unverändert, die Ausgabe ist `3`.
