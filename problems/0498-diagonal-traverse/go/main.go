package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // [1,2,4,7,5,3,6,8,9]
	fmt.Println(findDiagonalOrder([][]int{{1, 2}, {3, 4}}))                  // [1,2,3,4]
}

// Solution
func findDiagonalOrder(mat [][]int) []int {
	// Dimensions of the matrix
	m := len(mat)
	n := len(mat[0])

	// Solution slice with given size, as this is marginally faster
	solution := make([]int, m*n)

	// Index simply for adding to the solution array
	// row/column are central for iterating the array and are the end condition
	// up decides the direction
	for i, row, column, up := 0, 0, 0, true; row < m && column < n; i++ {
		// Add the current cell to the solution
		solution[i] = mat[row][column]

		// If moving upwards, we check whether we cross boundaries in...
		if up {
			up = false
			// ...the corner...
			if row == 0 && column == n-1 {
				row++
				// ...only the top edge...
			} else if row == 0 {
				column++
				// ...or the right edge
			} else if column == n-1 {
				row++
				// Otherwise, we simply move
			} else {
				row--
				column++
				up = true
			}
			// Moving downwards the same principle applies with bottom and left edges
		} else {
			up = true
			if row == m-1 && column == 0 {
				column++
			} else if row == m-1 {
				column++
			} else if column == 0 {
				row++
			} else {
				row++
				column--
				up = false
			}
		}
	}

	return solution
}
