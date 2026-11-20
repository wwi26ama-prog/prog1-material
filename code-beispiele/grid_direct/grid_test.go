package grid_direct

import "fmt"

func ExampleNew() {
	grid := New(3, 4)

	fmt.Println(grid)

	// Output:
	// [[. . . .] [. . . .] [. . . .]]
}

func ExampleSet() {
	grid := New(3, 4)
	Set(grid, 0, 0, "A")
	Set(grid, 1, 1, "B")
	Set(grid, 2, 2, "C")

	fmt.Println(grid)

	// Output:
	// [[A . . .] [. B . .] [. . C .]]
}

func ExampleGet() {
	grid := New(3, 4)
	Set(grid, 0, 0, "A")
	Set(grid, 1, 1, "B")
	Set(grid, 2, 2, "C")

	fmt.Println(Get(grid, 0, 0))
	fmt.Println(Get(grid, 1, 1))
	fmt.Println(Get(grid, 2, 2))
	fmt.Println(Get(grid, 0, 1)) // This will print the default value (Empty)

	// Output:
	// A
	// B
	// C
	// .
}

func ExamplePrint() {
	grid := New(3, 4)
	Set(grid, 0, 0, "A")
	Set(grid, 1, 1, "B")
	Set(grid, 2, 2, "C")

	Print(grid)

	// Output:
	// A | . | . | .
	// . | B | . | .
	// . | . | C | .
}
