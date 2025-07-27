package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0 1]
}

// Solution
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, iVal := range nums {
		// Calculate which other value would be necessary...
		diff := target - iVal
		// ...and then check whether it exists
		if j, ok := seen[diff]; ok {
			return []int{j, i}
		} else {
			seen[iVal] = i
		}
	}
	panic("No solution")
}

func twoSum_bruteForce(nums []int, target int) []int {
	for i, iVal := range nums {
		if iVal > target {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			jVal := nums[j]
			if (iVal + jVal) == target {
				return []int{i, j}
			}
		}
	}
	panic("No solution")
}
