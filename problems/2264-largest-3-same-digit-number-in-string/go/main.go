package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestGoodInteger("6777133339")) // "777"
	fmt.Println(largestGoodInteger("2300019"))    // "000"
	fmt.Println(largestGoodInteger("42352338"))   // ""
}

// Solution
func largestGoodInteger(num string) string {
	// Initialize for non-numeric character
	maxChar := '0' - 1

	lastChar := int32(num[0])
	nrSameChar := 0
	// Iterate the string
	for _, char := range num {
		// Count the occurrences of the same character
		if char == lastChar {
			nrSameChar++
			// If it is another character, reset
		} else {
			nrSameChar = 1
			lastChar = char
		}
		// If we count 3 same characters and the character is larger than the current maximum, overwrite
		if nrSameChar == 3 && char > maxChar {
			maxChar = char
		}
	}

	// If numeric character, generate string
	if maxChar >= '0' {
		return string([]rune{maxChar, maxChar, maxChar})
	}
	return ""
}
