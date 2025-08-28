package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortMatrix([][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}})) // [[8,2,3],[9,6,7],[4,5,1]]
	fmt.Println(sortMatrix([][]int{{0, 1}, {1, 2}}))                  // [[2,1],[1,0]]
}

// Solution
func sortMatrix(grid [][]int) [][]int {
	// Dimensions of the matrix
	n := len(grid)

	if n == 1 {
		return grid
	}

	// Lower Diagonals
	for i := 1; i <= n; i++ {
		dia := make([]int, i)

		// Extract the diagonal
		for row, col, j := n-i, 0, 0; j < i; row, col, j = row+1, col+1, j+1 {
			dia[j] = grid[row][col]
		}

		// Sort it descendingly
		sort.Slice(dia, func(i, j int) bool {
			return dia[i] > dia[j]
		})

		// Write it back into the matrix
		for row, col, j := n-i, 0, 0; j < i; row, col, j = row+1, col+1, j+1 {
			grid[row][col] = dia[j]
		}
	}
	// Upper Diagonals
	for i := n - 1; i >= 1; i-- {
		dia := make([]int, i)

		// Extract the diagonal
		for row, col, j := 0, n-i, 0; j < i; row, col, j = row+1, col+1, j+1 {
			dia[j] = grid[row][col]
		}

		// Sort it ascendingly
		sort.Slice(dia, func(i, j int) bool {
			return dia[i] < dia[j]
		})

		// Write it back into the matrix
		for row, col, j := 0, n-i, 0; j < i; row, col, j = row+1, col+1, j+1 {
			grid[row][col] = dia[j]
		}
	}

	return grid
}
