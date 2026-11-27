package card

import (
	"2026_11_27/karten/rank"
	"2026_11_27/karten/suit"
	"fmt"
)

func ExampleCard_Matches() {
	c1 := New(rank.Zehn, suit.Herz)
	c2 := New(rank.Ass, suit.Herz)
	c3 := New(rank.Zehn, suit.Karo)

	fmt.Println(c1.Matches(c2)) // gleiche Farbe
	fmt.Println(c1.Matches(c3)) // gleicher Rang
	fmt.Println(c2.Matches(c3)) // keine Übereinstimmung

	// Output:
	// true
	// true
	// false
}

func ExampleNew() {
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
