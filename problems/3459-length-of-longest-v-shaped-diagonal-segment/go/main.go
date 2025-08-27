package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(lenOfVDiagonal([][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}})) // 5
	fmt.Println(lenOfVDiagonal([][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}})) // 4
	fmt.Println(lenOfVDiagonal([][]int{{1, 2, 2, 2, 2}, {2, 2, 2, 2, 0}, {2, 0, 0, 0, 0}, {0, 0, 2, 2, 2}, {2, 0, 0, 2, 0}})) // 5
	fmt.Println(lenOfVDiagonal([][]int{{1}}))                                                                                 // 1
}

// Solution
func lenOfVDiagonal(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	maxPosLen := max(n, m)

	maxLen := 0

	// Iterate over all cells to find a starting point 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cell := grid[i][j]

			if cell != 1 {
				continue
			}

			lens := make([]int, 0)

			// Try to move into every direction
			for direction := range 4 {
				curLen := 1
				nextSequence := 2

				// Continuously move into a direction so long as it fits the sequence
				ii, jj, nextCell := moveDiagonal(grid, i, j, n, m, direction)
				for nextCell == nextSequence {
					nextSequence = switchNextSequence(nextSequence)
					curLen++
					// At each step, try to make the clockwise turn and calculate the result it would yield
					lens = append(lens, curLen+moveUntilEnd(grid, ii, jj, n, m, (direction+1)%4, nextSequence))
					ii, jj, nextCell = moveDiagonal(grid, ii, jj, n, m, direction)
				}
				// Add the result of no diverging paths
				lens = append(lens, curLen)
			}

			// Select the longest of all the possible paths
			maxCurLen := slices.Max(lens)
			if maxCurLen > maxLen {
				maxLen = maxCurLen
			}

			// If the result is already the longest it could be, don't bother analyzing anything else
			if maxLen == maxPosLen {
				return maxLen
			}
		}
	}

	return maxLen
}

// Simply move in a given direction according to the sequence without turns
func moveUntilEnd(grid [][]int, i, j, n, m, direction, nextSequence int) int {
	curLen := 0
	ii, jj, nextCell := moveDiagonal(grid, i, j, n, m, direction)
	for nextCell == nextSequence {
		nextSequence = switchNextSequence(nextSequence)
		curLen++
		ii, jj, nextCell = moveDiagonal(grid, ii, jj, n, m, direction)
	}
	return curLen
}

// Calculate the diagonal movement, i.e. the new position and its cell value
func moveDiagonal(grid [][]int, i, j, n, m, direction int) (int, int, int) {
	switch direction {
	case 0:
		nextI := i - 1
		nextJ := j - 1
		if nextI < 0 || nextJ < 0 {
			return nextI, nextJ, -1
		}
		return nextI, nextJ, grid[nextI][nextJ]
	case 1:
		nextI := i - 1
		nextJ := j + 1
		if nextI < 0 || nextJ >= m {
			return nextI, nextJ, -1
		}
		return nextI, nextJ, grid[nextI][nextJ]
	case 2:
		nextI := i + 1
		nextJ := j + 1
		if nextI >= n || nextJ >= m {
			return nextI, nextJ, -1
		}
		return nextI, nextJ, grid[nextI][nextJ]
	case 3:
		nextI := i + 1
		nextJ := j - 1
		if nextI >= n || nextJ < 0 {
			return nextI, nextJ, -1
		}
		return nextI, nextJ, grid[nextI][nextJ]
	}
	return i, j, -1
}

// Generate the next value in the sequence
func switchNextSequence(sequence int) int {
	if sequence == 2 {
		return 0
	} else {
		return 2
	}
}
