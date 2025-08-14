package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxTotalFruits([][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4))                          // 9
	fmt.Println(maxTotalFruits([][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 4)) // 14
	fmt.Println(maxTotalFruits([][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2))                          // 0
}

// Solution
func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	sum, maxSum := 0, 0

	// We use a sliding window
	for left, right := 0, 0; right < len(fruits); right++ {
		// Once an element enters the sliding window, it is added to the sum
		sum += fruits[right][1]
		// If our sliding window covers too much distance, we need to shrink it and remove the
		// elements that now fall outside from the sum
		for left <= right && minSteps(fruits[left][0], fruits[right][0], startPos) > k {
			sum -= fruits[left][1]
			left++
		}
		// Keep track of the largest sum at any point
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

// Calculate the distance expressed by the sliding window
func minSteps(left, right, start int) int {
	// Either go left first...
	a := abs(start-left) + (right - left)
	// ...or right first...
	b := abs(start-right) + (right - left)
	// ...and use the shorter way
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxTotalFruits_maps(fruits [][]int, startPos int, k int) int {
	fruitsAtPos := make(map[int]int)
	for _, entry := range fruits {
		fruitsAtPos[entry[0]] = entry[1]
	}

	leftPrefix := make([]int, k+1)
	for i := range k {
		leftPrefix[i+1] = leftPrefix[i] + fruitsAtPos[startPos-(i+1)]
	}

	rightPrefix := make([]int, k+1)
	for i := range k {
		rightPrefix[i+1] = rightPrefix[i] + fruitsAtPos[startPos+(i+1)]
	}

	res := 0
	for x := 0; x <= k; x++ {
		lSum := fruitsAtPos[startPos] + leftPrefix[x]
		if moveRight := k - 2*x; moveRight >= 0 {
			lSum += rightPrefix[moveRight]
		}
		if lSum > res {
			res = lSum
		}
		rSum := fruitsAtPos[startPos] + rightPrefix[x]
		if moveLeft := k - 2*x; moveLeft >= 0 {
			rSum += leftPrefix[moveLeft]
		}
		if rSum > res {
			res = rSum
		}
	}

	return res
}
