package moretypes

import "fmt"

func ExampleTime() {

	t1 := Time{13, 23}
	t2 := Time{Minutes: 30, Hours: 16}

	fmt.Println(t1.String())
	fmt.Println(t2)

	// Output:
	// 13:23
	// 16:30
}
