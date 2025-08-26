package main

import (
	"fmt"
)

func main() {
	fmt.Println(areaOfMaxDiagonal([][]int{{9, 3}, {8, 6}})) // 48
	fmt.Println(areaOfMaxDiagonal([][]int{{3, 4}, {4, 3}})) // 12
}

// Solution
func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDia, maxArea := 0, 0

	// Iterate all options
	for _, d := range dimensions {
		// Taking the square root is not necessary, as it has no impact on the ordering and the actual diagonal is not needed
		dia := d[0]*d[0] + d[1]*d[1]
		// If we found a longer diagonal, replace the area
		if dia > maxDia {
			maxDia = dia
			area := d[0] * d[1]
			maxArea = area
			// If the diagonal is of the same length, take the larger area
		} else if dia == maxDia {
			area := d[0] * d[1]
			maxArea = max(maxArea, area)
		}
	}

	return maxArea
}
