package moretypes

import "fmt"

func ExampleTime() {

	t1 := Time{13, 23}
	t2 := Time{Minutes: 30, Hours: 16}
	t3 := Time{Hours: 15, Minutes: 90}

	fmt.Println(t1.String())
	fmt.Println(t2)
	fmt.Println(t3)

	fmt.Println(t1 == t2)

	// Output:
	// 13:23
	// 16:30
	// 16:30
	// false
}

func ExampleTime_Normalize() {

	t1 := Time{Hours: 15, Minutes: 90}

	fmt.Println(t1)
	t1.Normalize()
	fmt.Println(t1)

	// Output:
	// 15:90
	// 16:30
}
