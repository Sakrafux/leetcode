package main

import (
	"fmt"
)

func main() {
	fmt.Println(findClosest(2, 7, 4)) // 1
	fmt.Println(findClosest(2, 5, 6)) // 2
	fmt.Println(findClosest(1, 5, 3)) // 0
}

// Solution
func findClosest(x int, y int, z int) int {
	// Absolute difference for x
	a := z - x
	if a < 0 {
		a *= -1
	}

	// Absolute difference for y
	b := z - y
	if b < 0 {
		b *= -1
	}

	// Return value based on absolute differences
	switch true {
	case a < b:
		return 1
	case a > b:
		return 2
	default:
		return 0
	}
}
