package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(sortVowels("lEetcOde")) // "lEOtcede"
	fmt.Println(sortVowels("lYmpH"))    // "lYmpH"
}

// Solution
func sortVowels(s string) string {
	// Copy the original string
	t := []byte(s)

	// Use slices to keep track of the vowels and their positions
	vowelLoc := make([]int, 0)
	vowels := make([]byte, 0)

	// Extract each vowel and its position
	for i, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			vowelLoc = append(vowelLoc, i)
			vowels = append(vowels, byte(c))
		}
	}

	// Sort the vowels
	slices.Sort(vowels)

	// Insert the vowels at their positions
	for i, v := range vowels {
		t[vowelLoc[i]] = v
	}

	return string(t)
}
