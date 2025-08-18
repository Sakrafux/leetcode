package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(judgePoint24([]int{4, 1, 8, 7})) // true
	fmt.Println(judgePoint24([]int{1, 2, 1, 2})) // false
	fmt.Println(judgePoint24([]int{1, 3, 4, 6})) // false
}

// Solution
func judgePoint24(cards []int) bool {
	nums := make([]float64, len(cards))
	for i := 0; i < len(cards); i++ {
		nums[i] = float64(cards[i])
	}
	return dfs(nums)
}

// EPS - Epsilon for numeric error tolerance
const EPS = 1e-6

func dfs(nums []float64) bool {
	// If we only have one element remaining, we combined all numbers somehow and only need to check, whether it is 24
	if len(nums) == 1 {
		return math.Abs(nums[0]-24.0) < EPS
	}

	// Select two numbers...
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// ...and remove them from consideration
			next := make([]float64, 0)
			for k := 0; k < len(nums); k++ {
				if i != k && j != k {
					next = append(next, nums[k])
				}
			}

			// Construct all possible calculations of those two numbers
			a, b := nums[i], nums[j]
			operations := []float64{a + b, a - b, b - a, a * b}
			if math.Abs(a) > EPS {
				operations = append(operations, b/a)
			}
			if math.Abs(b) > EPS {
				operations = append(operations, a/b)
			}

			// And add a single calculation back into consideration for the next recursion
			for _, op := range operations {
				if dfs(append(next, op)) {
					return true
				}
			}
		}
	}

	return false
}
