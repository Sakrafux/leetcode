package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))        // 321
	fmt.Println(reverse(-123))       // -321
	fmt.Println(reverse(120))        // 21
	fmt.Println(reverse(2147483647)) // 0
	fmt.Println(reverse(2147483642)) // 0
	fmt.Println(reverse(2147483602)) // 2063847412
}

// Solution
var UPPER_LIMIT int = math.MaxInt32 / 10
var LOWER_LIMIT int = math.MinInt32 / 10

func reverse(x int) int {
	res := 0
	// Continue until the entire number is processed
	for x != 0 {
		// Extract the digit
		digit := x % 10

		// If the result is higher/lower than the limit,
		// multiplying by 10 would lead to an over-/underflow
		if res > UPPER_LIMIT || res < LOWER_LIMIT {
			return 0
		}

		// Shift the result number left and add the digit
		res = 10*res + digit
		// Shift the working number right
		x /= 10
	}

	return res
}
