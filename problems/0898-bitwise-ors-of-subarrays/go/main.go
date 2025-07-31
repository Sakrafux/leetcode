package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarrayBitwiseORs([]int{0}))       // 1
	fmt.Println(subarrayBitwiseORs([]int{1, 1, 2})) // 3
	fmt.Println(subarrayBitwiseORs([]int{1, 2, 4})) // 6

}

// Solution
func subarrayBitwiseORs(arr []int) int {
	// A set keeping all solutions
	resultOrs := make(map[int]bool)

	// A set representing the OR values when adding x
	currentOrs := make(map[int]bool)
	for _, x := range arr {
		// Add the current number
		// Also represents starting a new subarray
		nextOrs := map[int]bool{x: true}
		resultOrs[x] = true
		// Try adding x to the previous subarrays
		for y := range currentOrs {
			or := x | y
			nextOrs[or] = true
			resultOrs[or] = true
		}
		// Update the current set to the one with x added
		currentOrs = nextOrs
	}

	return len(resultOrs)
}

func subarrayBitwiseORs_bruteForce(arr []int) int {
	n := len(arr)

	uniqVals := make(map[int]bool)

	res := 0

	for i := range arr {
		orVal := 0
		for j := i; j < n; j++ {
			orVal |= arr[j]
			if val := uniqVals[orVal]; !val {
				uniqVals[orVal] = true
				res++
			}
		}
	}

	return res
}
