package suit

// Suit repräsentiert die Farbe einer Spielkarte.
type Suit string

const (
	Karo  Suit = "♦"
	Herz  Suit = "❤"
	Pik   Suit = "♠"
	Kreuz Suit = "♣"
)

func (s Suit) String() string {
	return string(s)
}
