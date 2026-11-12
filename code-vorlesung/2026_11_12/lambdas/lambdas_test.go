package lambdas

import (
	"fmt"
)

func Example() {
	x := 23

	Foo := func() {
		x = 42
		fmt.Println(x)
	}

	Foo()
	fmt.Println(x)

	// Output:
}

func ExampleSort() {
	n := []int{1, 5, 2, 4, 3}

	foo := func(x, y int) bool {
		return x > y
	}

	fmt.Println(n)
	Sort(n, foo)
	fmt.Println(n)

	// Output:
}

func Sort(numbers []int, cmp func(int, int) bool) {
	for range numbers {
		for j := range numbers[:len(numbers)-1] {
			if cmp(numbers[j], numbers[j+1]) {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}
