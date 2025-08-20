package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}})) // 15
	fmt.Println(countSquares([][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}))          // 7
}

// Solution
func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	sum := 0

	// Iterate each cell of the matrix
	for i := range m {
		for j := range n {
			// If we are not at the top or left edge, our current cell could be the end of a square
			if i > 0 && j > 0 && matrix[i][j] == 1 {
				// Get the 3 immediate top and/or left neighbors
				left := matrix[i][j-1]
				top := matrix[i-1][j]
				topleft := matrix[i-1][j-1]

				// The minimum of the neighbors is the basis for the possible squares at this cell
				matrix[i][j] = min(left, top, topleft) + 1
			}
			sum += matrix[i][j]
		}
	}

	return sum
}
