package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numberOfPairs([][]int{{1, 1}, {2, 2}, {3, 3}})) // 0
	fmt.Println(numberOfPairs([][]int{{6, 2}, {4, 4}, {2, 6}})) // 2
	fmt.Println(numberOfPairs([][]int{{3, 1}, {1, 3}, {1, 1}})) // 2
	fmt.Println(numberOfPairs([][]int{{0, 1}, {1, 3}, {6, 1}})) // 2
	fmt.Println(numberOfPairs([][]int{{1, 5}, {2, 0}, {5, 5}})) // 2
	fmt.Println(numberOfPairs([][]int{{1, 6}, {0, 9}, {0, 3}})) // 2
}

// Solution
func numberOfPairs(points [][]int) int {
	// Sort the pairs from top-left to bottom-right, with left taking precedence over top
	sort.Slice(points, func(i, j int) bool {
		pair1, pair2 := points[i], points[j]
		// If an element is further left, it ranks first
		if pair1[0] < pair2[0] {
			return true
		}
		// Only if 2 elements are at the same x, rank the higher first
		if pair1[0] == pair2[0] {
			return pair1[1] > pair2[1]
		}

		return false
	})

	res := 0
	// Iterate over all points
	for i := 0; i < len(points); i++ {
		// Use a boundary variable to ensure no points are inside a rectangle
		bx, by := -1, -1
		// Iterate over all later points for checking combinations
		for j := i + 1; j < len(points); j++ {
			pair1, pair2 := points[i], points[j]
			// Check that the first point is to the top-left and the second point doesn't include a point between
			// by checking against the boundary
			if (pair1[0] <= pair2[0] && pair1[1] >= pair2[1]) && (pair2[0] > bx && pair2[1] > by) {
				res++
				// Update the boundary, which can only be a strict monotonic increase
				bx = points[j][0]
				by = points[j][1]
			}
		}
	}

	return res
}
