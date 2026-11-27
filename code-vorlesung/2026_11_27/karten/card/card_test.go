package card

import (
	"2026_11_27/karten/rank"
	"2026_11_27/karten/suit"
	"fmt"
)

func ExampleNew() {
	// Beispielhafte Verwendung der Card-Struktur
	c1 := New(rank.Zehn, suit.Herz)
	c2 := New(rank.Ass, suit.Herz)

	fmt.Println(c1)
	fmt.Println(c2)

	// Output:
	// ┌─────────┐
	// │ 10      │
	// │         │
	// │    ❤    │
	// │         │
	// │      10 │
	// └─────────┘
	// ┌─────────┐
	// │ A       │
	// │         │
	// │    ❤    │
	// │         │
	// │       A │
	// └─────────┘
}
