package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations([][]int{{1, 2}, {2, 4}})) // 3
	fmt.Println(minOperations([][]int{{2, 6}}))         // 4
	fmt.Println(minOperations([][]int{{3, 10}}))        // 8
	fmt.Println(minOperations([][]int{{3, 7}, {3, 7}})) // 10
}

// Solution
func minOperations(queries [][]int) int64 {
	var ans int64 = 0

	// Iterate the queries
	for _, q := range queries {
		l, r := q[0], q[1]

		// Total Single Element Reductions, i.e. how many reduction operations in total
		var tser int64 = 0

		// Iterate divisible blocks as defined by their bit length, i.e. all elements in those blocks need the same
		// number of operations to reach 0
		for k := 1; k <= 31; k++ {
			// Lower boundary 2^{k-1}
			low := int64(1) << (k - 1)
			// Upper boundary 2^k-1
			high := (int64(1) << k) - 1
			// The range of the block is greater than the query range, so we are finished for the query
			if low > int64(r) {
				break
			}
			// Correct to the edges of our query range
			a := max(int64(l), low)
			b := min(int64(r), high)
			// If the implied block range is smaller than the beginning of our query range, continue with the next block
			if a > b {
				continue
			}
			// Amount of elements in the block/range
			cnt := b - a + 1
			// Necessary operations to reduce the elements of this block to 0
			stepsForK := (k + 1) / 2
			// Add to total sum of operations in this block/range
			tser += cnt * int64(stepsForK)
		}

		// Since our operations affect 2 elements at once, divide the total number of operations by 2,
		// while considering odd numbers
		ops := (tser + 1) / 2
		ans += ops
	}

	return ans
}
