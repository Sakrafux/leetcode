package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // [[2,2,2,2]]
}

// Solution
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	// Handle special case of too small array
	if n < 4 {
		return [][]int{}
	}

	// Sort the array so the order of the elements is assured
	slices.Sort(nums)

	res := make([][]int, 0)

	// Fix one number with an loop
	for i, num1 := range nums {
		// Prevent duplicate results
		if i > 0 && nums[i-1] == num1 {
			continue
		}

		// Fix a second number with another loop
		for j := i + 1; j < n; j++ {
			num2 := nums[j]
			// Prevent duplicate results
			if j > i+1 && nums[j-1] == num2 {
				continue
			}

			// Set 2 pointers to the edges of the remaining array space
			for l, r := j+1, n-1; l < r; {
				num3, num4 := nums[l], nums[r]
				sum := num1 + num2 + num3 + num4

				// Is the sum smaller than the target, increase the sum
				if sum < target {
					l++
					// If it is larger, decrease it
				} else if sum > target {
					r--
					// Otherwise, we found a result
				} else {
					res = append(res, []int{num1, num2, num3, num4})
					l++

					// Prevent duplicate results
					for num3 == nums[l] && l < r {
						l++
					}
				}
			}
		}
	}

	return res
}
