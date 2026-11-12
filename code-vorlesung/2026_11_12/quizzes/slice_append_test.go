package quizzes

func ExampleCallAddZero() {
	CallAddZero()

	// Output:
	// [2 3 5]
}

// - (15) `primes` ist ein Slice der Länge `3` mit den Zahlen `2`, `3`
//   und `5`. Das Array dahinter hat keinen Platz für weitere Zahlen.
// - (9, 16) `numbers` ist eine Kopie des Slices `primes`: mit derselben
//   Länge und demselben Array.
// - (10) `append` gibt einen neuen Slice der Länge `4` zurück, mit einem
//   neuen, größeren Array. Die Zuweisung ändert nur die Kopie `numbers`.
// - (17) `primes` ist unverändert, die Ausgabe ist `[2 3 5]`.
// - (16) Damit die `0` ankommt, gibt man den Slice zurück, wie `append`
//   selbst: `primes = AddZero(primes)`.
