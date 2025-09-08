package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNoZeroIntegers(2))    // [1,1]
	fmt.Println(getNoZeroIntegers(11))   // [2,9]
	fmt.Println(getNoZeroIntegers(1057)) // [58,999]
}

// Solution
func getNoZeroIntegers(n int) []int {
	// Start with non-zero a
	a, b := 1, n-1
	// Add/Subtract variables by 1 until no zero digits remain
	for ; checkHasO(a) || checkHasO(b); a, b = a+1, b-1 {
	}
	return []int{a, b}
}

// Check for zero digit
func checkHasO(n int) bool {
	for ; n > 0; n /= 10 {
		if n%10 == 0 {
			return true
		}
	}
	return false
}
