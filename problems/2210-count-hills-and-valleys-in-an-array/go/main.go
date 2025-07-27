package main

import "fmt"

func main() {
	fmt.Println(countHillValley([]int{2, 4, 1, 1, 6, 5})) // 3
	fmt.Println(countHillValley([]int{6, 6, 5, 5, 4, 1})) // 0
}

// Solution
func countHillValley(nums []int) int {
	count := 0
	l := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		m := nums[i]
		r := nums[i+1]
		if l == m || r == m {
			continue
		}
		if (m > l && m > r) || (m < l && m < r) {
			count++
		}
		l = m
	}
	return count
}
