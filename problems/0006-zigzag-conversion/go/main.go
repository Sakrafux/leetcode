package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) // PINALSIGYAHRPI
	fmt.Println(convert("A", 1))              // A
	fmt.Println(convert("AB", 1))             // AB
}

// Solution
func convert(s string, numRows int) string {
	// Special case of 1 row needs no conversion
	if numRows == 1 {
		return s
	}

	// Keep every row as its own slice
	rows := make([][]byte, numRows)
	for i := range numRows {
		rows[i] = make([]byte, 0)
	}

	// Iterate through the string and keeping track of the current row index
	// For zig-zag this index simply needs to increase and decrease alternately
	for i, rowIndex, step := 0, 0, -1; i < len(s); i, rowIndex = i+1, rowIndex+step {
		// If we reached the end (or beginning), change direction
		if rowIndex == 0 || rowIndex == numRows-1 {
			step *= -1
		}

		// Append to the corresponding row
		rows[rowIndex] = append(rows[rowIndex], s[i])
	}

	// Build the result by combining all rows
	res := make([]byte, 0)
	for _, row := range rows {
		res = append(res, row...)
	}

	return string(res)
}
