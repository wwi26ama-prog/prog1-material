package bibliothek

import (
	"time"
)

type Ausleihe struct {
	Kundennummer string
	ISBN         string
	Exemplar     int
	Auhsleihzeit time.Time
	Dauer        time.Duration
}
