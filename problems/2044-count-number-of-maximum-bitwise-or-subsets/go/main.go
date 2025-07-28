package main

import (
	"fmt"
)

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))       // 2
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))    // 7
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5})) // 6
}

// Solution
func countMaxOrSubsets(nums []int) int {
	// Reserve space for all possible numbers as given by the constraints
	occurences := make([]int, 1<<17)
	// Add default for empty subset
	occurences[0] = 1
	maxValue := 0

	// Iterate through all numbers
	for _, num := range nums {
		// Check going from the current maximum value backwards...
		for i := maxValue; i >= 0; i-- {
			// ...and the combination with the new number inherits
			// all occurences of its previous subsets
			occurences[i|num] += occurences[i]
		}
		// Update the maximum value accordingly
		maxValue |= num
	}

	return occurences[maxValue]
}

func countMaxOrSubsets_bruteForce(nums []int) int {
	n := len(nums)
	subsetSumsCounts := make(map[int]int)
	maxSum := 0

	// All subset permutations are 2^n
	for i := 0; i < (1 << n); i++ {
		sum := 0
		// Transform each array index into a bitmask and check whether it has
		// any overlap with (i.e. is part of) the current iteration
		for j := range n {
			if i&(1<<j) != 0 {
				sum |= nums[j]
			}
		}
		subsetSumsCounts[sum] += 1
		if sum > maxSum {
			maxSum = sum
		}
	}

	return subsetSumsCounts[maxSum]
}
