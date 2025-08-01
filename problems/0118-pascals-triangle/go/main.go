package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
	fmt.Println(generate(1)) // [[1]]
}

// Solution
func generate(numRows int) [][]int {
	// Prepare the result array
	res := make([][]int, numRows)
	// The first row is special
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		// Take the previous row
		prevRow := res[i-1]
		// Create the current row one larger than the previous
		row := make([]int, i+1)
		// Calculate each number in the new row
		for j := 0; j < i+1; j++ {
			// The first number is to the left in the previous row or 0, if out of bounds
			num1 := 0
			if j > 0 {
				num1 = prevRow[j-1]
			}
			// The second number is at the same position in the previous row or 0, if out of bounds
			num2 := 0
			if j < i {
				num2 = prevRow[j]
			}

			row[j] = num1 + num2
		}
		res[i] = row
	}

	return res
}
