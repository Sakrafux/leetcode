package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{1, 1}))                      // 1
}

// Solution
func maxArea(height []int) int {
	area := 0

	// Keep two pointers at the array boundaries and move them inwards
	for l, r := 0, len(height)-1; l < r; {
		lHeight := height[l]
		rHeight := height[r]

		// Calculate the area with the lower height and move the lower pointer inwards
		newArea := 0
		if lHeight < rHeight {
			newArea = (r - l) * lHeight
			l++
		} else {
			newArea = (r - l) * rHeight
			r--
		}

		// Check every calculated area against the current maximum
		if newArea > area {
			area = newArea
		}
	}

	return area
}
