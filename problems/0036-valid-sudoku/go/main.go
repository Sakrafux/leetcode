package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', ' ', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})) // true
	fmt.Println(isValidSudoku([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})) // false
}

// Solution
func isValidSudoku(board [][]byte) bool {
	// Map capturing the occurrences on columns
	columnMap := map[int]map[byte]bool{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}}
	// Map capturing the occurrences on 3x3 blocks
	blockMap := map[int]map[byte]bool{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}}

	for i, row := range board {
		// Map capturing the occurrences on a given row
		rowMap := make(map[byte]bool)
		for j, cell := range row {
			// Ignore blanks
			if cell == '.' {
				continue
			}

			// Check for duplicates in the row...
			if _, ok := rowMap[cell]; ok {
				return false
			}
			rowMap[cell] = true

			// ...then in the block...
			blockIndex := (j / 3) + (i/3)*3
			block := blockMap[blockIndex]
			if _, ok := block[cell]; ok {
				return false
			}
			block[cell] = true

			// ...and finally in the columns
			if _, ok := columnMap[j][cell]; ok {
				return false
			}
			columnMap[j][cell] = true
		}
	}

	// If no duplicates occurred until the end, then the board is valid
	return true
}
