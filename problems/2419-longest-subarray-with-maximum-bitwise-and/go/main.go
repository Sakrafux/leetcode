package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 3, 2, 2}))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              // 2
	fmt.Println(longestSubarray([]int{1, 2, 3, 4}))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    // 1
	fmt.Println(longestSubarray([]int{96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 96317, 279979}))                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         // 1
	fmt.Println(longestSubarray([]int{395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 153490, 330001, 330001, 330001, 330001, 330001, 330001, 330001, 37284, 470030, 470030, 470030, 470030, 470030, 470030, 156542, 226743})) // 24

}

// Solution
func longestSubarray(nums []int) int {
	maxAnd := nums[0]
	maxLength := 1

	// Keep track of the current segment
	currentLength := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// Make use of the fact the largest AND can only be the largest number
		if num > maxAnd {
			currentLength = 1
			maxLength = 1
			maxAnd = num
			// If we are on a continued stretch of the highest number, count the length
		} else if num == maxAnd && nums[i-1] == num {
			currentLength++
			// In case we are on a different stretch of the highest number, check it is longer
			if currentLength > maxLength {
				maxLength = currentLength
			}
			// Reset if we left the stretch
		} else {
			currentLength = 1
		}
	}

	return maxLength
}
