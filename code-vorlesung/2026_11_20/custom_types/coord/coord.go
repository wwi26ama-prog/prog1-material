package coord

import "2026_11_20/custom_types/length"

type Coord struct {
	Longitude float64
	Latitude  float64
}

func (c Coord) Distance(o Coord) length.Length {

	// TODO: Die Entfernung wirklich berechnen.

	return length.FromKilometers(42)
}
