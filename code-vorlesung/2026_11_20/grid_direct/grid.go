package grid_direct

import (
	"fmt"
	"strings"
)

const Empty = "."

// New erzeugt eine 2D-String-Slice mit der
// angegebenen Anzahl an Zeilen und Spalten.
func New(rows, cols int) [][]string {
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, cols)
		for j := range grid[i] {
			grid[i][j] = Empty
		}
	}
	return grid
}

// Set setzt den Wert an der angegebenen Position (row, col) im Grid.
// Falls die Position außerhalb der Grenzen liegt, wird eine Panik ausgelöst.
func Set(grid [][]string, row, col int, value string) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		panic(fmt.Sprintf("Index out of bounds (row=%d, col=%d) for grid with %d rows and %d cols", row, col, len(grid), len(grid[0])))
	}
	grid[row][col] = value
}

// Get gibt den Wert an der angegebenen Position (row, col) im Grid zurück.
// Falls die Position außerhalb der Grenzen liegt, wird eine Panik ausgelöst.
func Get(grid [][]string, row, col int) string {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		panic(fmt.Sprintf("Index out of bounds (row=%d, col=%d) for grid with %d rows and %d cols", row, col, len(grid), len(grid[0])))
	}
	return grid[row][col]
}

// Print gibt das Grid in einer lesbaren Form auf der Konsole aus.
func Print(grid [][]string) {
	for _, row := range grid {
		rowstring := strings.Join(row, " | ")
		fmt.Println(rowstring)
	}
}
