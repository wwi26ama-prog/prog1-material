package power

import "fmt"

// Power2 berechnet 2^n
func Power2(n int) int {
	if n == 0 {
		return 1
	}

	result := 2 * Power2(n-1)
	return result
}

func ExamplePower2() {
	x := Power2(5)
	fmt.Println(x)

	// Output:
	// 32
}
