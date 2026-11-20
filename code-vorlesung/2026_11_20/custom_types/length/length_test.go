package length

import "fmt"

func ExampleLength() {
	//var l Length = 5000000
	//l := FromMillimeters(5000000)
	l := FromKilometers(5)

	cm := l.Centimeters()
	m := l.Meters()
	km := l.Kilometers()

	fmt.Println(l)
	fmt.Println(cm)
	fmt.Println(m)
	fmt.Println(km)

	// Output:
	// 5000000
	// 500000
	// 5000
	// 5
}
