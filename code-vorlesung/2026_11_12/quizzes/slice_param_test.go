package quizzes

func ExampleCallDoubleAll() {
	CallDoubleAll()

	// Output:
	// [4 6 10]
}

// - (17) `primes` ist ein Slice, der auf ein Array mit den Zahlen
//   `2`, `3` und `5` verweist.
// - (9, 18) Auch hier bekommt der Parameter eine Kopie: `numbers` ist
//   eine Kopie des Slices `primes`, nicht des Arrays.
// - (10, 11) Beide Slices verweisen auf dasselbe Array. `DoubleAll`
//   ändert also die Zahlen, die auch `primes` sieht.
// - (19) Die Ausgabe zeigt deshalb die verdoppelten Zahlen `[4 6 10]`.
