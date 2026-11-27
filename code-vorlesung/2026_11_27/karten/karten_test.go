package karten

import (
	"2026_11_27/karten/rank"
	"2026_11_27/karten/suit"
	"fmt"
)

func Example() {
	kreuz := suit.Kreuz
	dame := rank.Dame

	fmt.Println(kreuz)
	fmt.Println(dame)

	// Output:
	// ♣
	// D
}
