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

Outer:
	for _, fruit := range fruits {
		for i, basket := range baskets {
			if fruit <= basket {
				baskets[i] = -1
				continue Outer
			}
		}
		unplaced++
	}

	return unplaced
}
