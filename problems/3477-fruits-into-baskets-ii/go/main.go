package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfUnplacedFruits([]int{4, 2, 5}, []int{3, 5, 4})) // 1
	fmt.Println(numOfUnplacedFruits([]int{3, 6, 1}, []int{6, 4, 7})) // 0
}

// Solution
func numOfUnplacedFruits(fruits []int, baskets []int) int {
	unplaced := 0

	// The outer loop tries to place each fruit
Outer:
	for _, fruit := range fruits {
		// The inner loop checks each basket
		for i, basket := range baskets {
			// If a basket is sufficient, use it by setting its capacity negative and continue with the next fruit
			if fruit <= basket {
				baskets[i] = -1
				continue Outer
			}
		}
		// If no basket was sufficient, it goes unplaced
		unplaced++
	}

	return unplaced
}
