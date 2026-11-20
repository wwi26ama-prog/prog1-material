package duration

import "fmt"

func ExampleDuration() {
	d1 := FromSeconds(42)
	d2 := FromMinutes(1)
	d3 := FromHours(1)

	// Die drei Werte ausgeben:
	fmt.Println(d1.Seconds())
	fmt.Println(d2.Seconds())
	fmt.Println(d3.Seconds())

	// Output:
	// 42
	// 60
	// 3600
}
