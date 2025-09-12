package main

import (
	"fmt"
)

func main() {
	fmt.Println(doesAliceWin("leetcoder")) // true
	fmt.Println(doesAliceWin("bbcd"))      // false
}

// Solution
func doesAliceWin(s string) bool {
	// Check if any vowel exists
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			return true
		}
	}

	return false
}
