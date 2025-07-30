package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      // 0

}

// Solution
func threeSumClosest(nums []int, target int) int {
	// Sort the numbers so we are assured of the ordering
	slices.Sort(nums)

	// Initialize the result to some impossibly high value (given the constraints)
	res := 100000

	// Fix one number
	for i, num1 := range nums {
		// Use a left and right pointer for the remaining array space
		for l, r := i+1, len(nums)-1; l < r; {
			num2, num3 := nums[l], nums[r]
			// Get the sum of the chosen numbers...
			sum := num1 + num2 + num3

			// ...and check whether that sum is closer to the target than the current best
			if int(math.Abs(float64(target-sum))) < int(math.Abs(float64(target-res))) {
				// If so, it is the current best
				res = sum
			}

			// If our sum is higher than the target, then we must decrease it by moving the
			// right pointer leftwards (as our array is sorted, this is assured to be <= to before)
			if sum > target {
				r--
				// Similarly, in the other direction we need to increase the left pointer
			} else if sum < target {
				l++
				// If we reached the target, we can simply stop
			} else {
				return res
			}
		}
	}

	return res
}
