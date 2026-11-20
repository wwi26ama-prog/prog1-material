package main

import (
	"2026_11_20/grid_direct"
	"fmt"
)

type ShotGrid struct {
	Grid [][]string
}

func (s ShotGrid) Print() {
	grid_direct.Print(s.Grid)
}

// Dieses Programm demonstriert die Verwendung der Funktionen aus dem Paket
// "grid_direct" anhand einer vereinfachten Variante des Spiels "Schiffe versenken".
//
// Es erstellt zwei 10x10 Grids:
// Das erste Grid repräsentiert das Spielfeld eines Spielers, während das zweite Grid
// die Sicht dieses Spielers auf das gegnerische Spielfeld darstellt.
// D.h. auf dem ersten Grid werden einige Schiffe platziert, während das zweite
// Grid anzeigt, welche Positionen der Spieler bereits beschossen hat.
func main() {
	ships := grid_direct.New(10, 10)
	shots := ShotGrid{Grid: grid_direct.New(10, 10)}

	// Platziere ein Schiff der Länge 4 horizontal auf dem Spielfeld.
	SetFour(ships, 2, 3)

	// Platziere ein Schiff der Länge 3 vertikal auf dem Spielfeld.
	SetThree(ships, 5, 1)

	// Markiere einige Positionen als bereits beschossen.
	Hit(shots, 2, 3)  // Treffer
	Hit(shots, 5, 1)  // Treffer
	Miss(shots, 0, 0) // Verfehlt
	Miss(shots, 9, 9) // Verfehlt

	// Zeige die beiden Grids an.
	fmt.Println("Shots:")
	shots.Print()
	fmt.Println()
	fmt.Println("Ships:")
	grid_direct.Print(ships)
}

func SetFour(grid [][]string, row, col int) {
	for i := range 4 {
		grid_direct.Set(grid, row, col+i, "S")
	}
}

func SetThree(grid [][]string, row, col int) {
	for i := range 3 {
		grid_direct.Set(grid, row+i, col, "S")
	}
}

func Hit(grid ShotGrid, row, col int) {
	grid_direct.Set(grid.Grid, row, col, "X") // Treffer
}

func Miss(grid ShotGrid, row, col int) {
	grid_direct.Set(grid.Grid, row, col, "O") // Verfehlt
}
