package card

import (
	"2026_11_27/karten/rank"
	"2026_11_27/karten/suit"
	"fmt"
	"strings"
)

type Card struct {
	Rank rank.Rank
	Suit suit.Suit
}

func New(rank rank.Rank, suit suit.Suit) Card {
	return Card{
		Rank: rank,
		Suit: suit,
	}
}

// String gibt eine Darstellung der Karte als AsciiArt zurück.
func (c Card) String() string {
	lines := []string{
		"┌─────────┐",
		fmt.Sprintf("│ %-2s      │", c.Rank.String()),
		"│         │",
		fmt.Sprintf("│    %s    │", c.Suit.String()),
		"│         │",
		fmt.Sprintf("│      %2s │", c.Rank.String()),
		"└─────────┘",
	}

	return strings.Join(lines, "\n")
}
