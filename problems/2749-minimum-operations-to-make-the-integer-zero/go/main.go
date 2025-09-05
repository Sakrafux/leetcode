package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(makeTheIntegerZero(3, -2)) // 3
	fmt.Println(makeTheIntegerZero(5, 7))  // -1
}

// Solution
func makeTheIntegerZero(num1 int, num2 int) int {
	// Iterate upon the number of operations until success (hard bounded by the input constraints)
	for k := 1; ; k++ {
		// Each operation necessarily carries the -num2, thus this expression calculates the number we need to reach
		// purely by 2^i operations
		x := num1 - num2*k
		// If we have more operations than the value to reach, it simply isn't possible
		if x < k {
			return -1
		}
		// Check if we have enough operations to service every bit that occurs
		if k >= bits.OnesCount(uint(x)) {
			return k
		}
	}
}
