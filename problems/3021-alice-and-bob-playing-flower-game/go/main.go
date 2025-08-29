package main

import (
	"fmt"
)

func main() {
	fmt.Println(flowerGame(3, 2)) // 3
	fmt.Println(flowerGame(1, 1)) // 0
	fmt.Println(flowerGame(4, 4)) // 8
}

// Solution
func flowerGame(n int, m int) int64 {
	res := int64(0)

	// Ensure that n is odd if possible
	if n%2 == 0 {
		n, m = m, n
	}

	// Calculate all pairs with odd x, recognizing n might be odd
	res += int64((n + 1) / 2 * m / 2)
	// Calculate all pairs with odd y
	res += int64(n / 2 * m / 2)

	return res
}
