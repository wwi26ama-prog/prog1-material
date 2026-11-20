package coord

import "fmt"

func ExampleCoord() {
	c1 := Coord{8.5, 49.5}
	c2 := Coord{0, 0}

	d1 := c1.Distance(c2)

	// Die beiden Koordinaten komplett ausgeben:
	fmt.Println(c1)
	fmt.Println(c2)

	// Nur den Längengrad von c1 ausgeben:
	fmt.Println(c1.Longitude)

	// Den Breitengrad von c2 verändern:
	c2.Latitude = 42
	fmt.Println(c2)

	// Die Distanz ausgeben:
	fmt.Println(d1)

	// Output:
}
