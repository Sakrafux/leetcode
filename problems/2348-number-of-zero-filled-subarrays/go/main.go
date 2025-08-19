package main

import (
	"fmt"
)

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4})) // 6
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))       // 9
	fmt.Println(zeroFilledSubarray([]int{2, 10, 2019}))            // 0
}

// Solution
func zeroFilledSubarray(nums []int) int64 {
	// Track the sum of subarrays
	var sum int64 = 0
	// Track the amount of consecutive zeros
	var consecZeros int64 = 0
	// Iterate all the numbers of the array
	for _, num := range nums {
		// If it is 0, all the previous zeros can create a new subarray ending at the current number
		// plus a new 1-element subarray
		if num == 0 {
			sum += 1 + consecZeros
			consecZeros++
			// Otherwise, the consecutive zeros are reset
		} else {
			consecZeros = 0
		}
	}

	return sum
}
