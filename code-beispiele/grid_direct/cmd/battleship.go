package main

import (
	"fmt"
	"grid_direct"
)

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
	shots := grid_direct.New(10, 10)

	// Platziere ein Schiff der Länge 4 horizontal auf dem Spielfeld.
	grid_direct.Set(ships, 2, 3, "S")
	grid_direct.Set(ships, 2, 4, "S")
	grid_direct.Set(ships, 2, 5, "S")
	grid_direct.Set(ships, 2, 6, "S")

	// Platziere ein Schiff der Länge 3 vertikal auf dem Spielfeld.
	grid_direct.Set(ships, 5, 1, "S")
	grid_direct.Set(ships, 6, 1, "S")
	grid_direct.Set(ships, 7, 1, "S")

	// Markiere einige Positionen als bereits beschossen.
	grid_direct.Set(shots, 2, 3, "X") // Treffer
	grid_direct.Set(shots, 5, 1, "X") // Treffer
	grid_direct.Set(shots, 0, 0, "O") // Verfehlt
	grid_direct.Set(shots, 9, 9, "O") // Verfehlt

	// Zeige die beiden Grids an.
	fmt.Println("Shots:")
	grid_direct.Print(shots)
	fmt.Println()
	fmt.Println("Ships:")
	grid_direct.Print(ships)
}
