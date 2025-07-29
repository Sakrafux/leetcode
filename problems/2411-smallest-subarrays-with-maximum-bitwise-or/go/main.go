package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestSubarrays([]int{1, 0, 2, 1, 3}))    // [3,3,2,2,1]
	fmt.Println(smallestSubarrays([]int{1, 2}))             // [2,1]
	fmt.Println(smallestSubarrays([]int{4, 0, 5, 6, 3, 2})) // [4,3,2,2,1,1]
}

// Solution
func smallestSubarrays(nums []int) []int {
	n := len(nums)
	lastSeen := make([]int, 30)
	res := make([]int, n)

	// Every result must be atleast of length 1
	for i := range res {
		res[i] = 1
	}

	// Iterate the numbers in reverse to automatically capture the last occurences of bits
	for i := n - 1; i >= 0; i-- {
		// Iterate all possible bits
		for j := range lastSeen {
			// Does the current number contain the bit
			if nums[i]&(1<<j) > 0 {
				lastSeen[j] = i
			}
			// Check the distance a bit has been seen before
			if alt := lastSeen[j] - i + 1; alt > res[i] {
				res[i] = alt
			}
		}
	}

	return res
}

func smallestSubarrays_maps(nums []int) []int {
	n := len(nums)

	maxOr := 0
	for _, num := range nums {
		maxOr |= num
	}

	bitMap := make(map[int][]int)

	for i, x := 0, maxOr; x > 0; i, x = i+1, x>>1 {
		if x%2 == 1 {
			bitMap[1<<i] = make([]int, n)
		}
	}

	for key, value := range bitMap {
		for i := n - 1; i >= 0; i-- {
			num := nums[i]
			if key&num != 0 {
				value[i] = 1
			} else if i+1 < n && value[i+1] != 0 {
				value[i] = value[i+1] + 1
			}
		}
	}

	res := make([]int, n)

	for i := range nums {
		maxLen := 1
		for _, value := range bitMap {
			if l := value[i]; l > maxLen {
				maxLen = l
			}
		}
		res[i] = maxLen
	}

	return res
}

func smallestSubarrays_bruteForce(nums []int) []int {
	maxOr := 0
	for _, num := range nums {
		maxOr |= num
	}

	res := make([]int, len(nums))

	for i, num := range nums {
		subMaxOr := num
		subArrSize := 1
		realSubArrSize := 1
		for j := i + 1; j < len(nums) && subMaxOr < maxOr; j++ {
			prevSubMaxOr := subMaxOr
			subMaxOr |= nums[j]
			subArrSize++
			if prevSubMaxOr != subMaxOr {
				realSubArrSize = subArrSize
			}
		}
		res[i] = realSubArrSize
	}

	return res
}
