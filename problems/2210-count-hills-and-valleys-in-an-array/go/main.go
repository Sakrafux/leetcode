package main

import "fmt"

func main() {
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5})) // 3
	fmt.Println(countHillValley([]int{6, 6, 5, 5, 4, 1})) // 0
}

// Solution
func countHillValley(nums []int) int {
	count := 0
	// The left value may be persistent over iterations and is thus defined outside
	l := nums[0]
	// We iterate by moving our cursor element through the array, except for the outer-most elements
	for i := 1; i < len(nums)-1; i++ {
		m := nums[i]
		r := nums[i+1]
		// If we don't have different values for neighbours, we just keep moving/iterating,
		// letting the left neighbour stay in place until we have found the end of the plateau
		if l == m || r == m {
			continue
		}
		// At the end of a plateau (which may be simply 1 element wide), we check against our rules
		if (m > l && m > r) || (m < l && m < r) {
			count++
		}
		// Now the left neighbour moves along
		l = m
	}
	return count
}
