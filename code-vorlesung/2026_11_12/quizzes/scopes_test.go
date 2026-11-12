package quizzes

func ExampleScopes() {
	Scopes()

	// Output:
	// 3
	// 3
	// 1
}

// - (10) `x := 1` deklariert im Block der Funktion eine Variable `x`.
// - (11, 13) Auch ohne `if` oder `for` öffnen geschweifte Klammern einen
//   Block mit eigenem Gültigkeitsbereich (Scope).
// - (12) `:=` deklariert eine zweite Variable `x`, die das äußere `x`
//   in diesem Block verdeckt (Shadowing).
// - (14) `=` ist eine Zuweisung, keine Deklaration: Sie ändert das `x`
//   aus Zeile 12 von `2` auf `3`.
// - (15, 17) Beide Ausgaben beziehen sich auf dieses `x` und geben `3` aus.
// - (19) Nach dem Block ab Zeile 11 ist wieder das `x` aus Zeile 10 sichtbar,
//   unverändert mit dem Wert `1`.
