package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(maxFreqSum("successes")) // 6
	fmt.Println(maxFreqSum("aeiaeia"))   // 3
}

// Solution
func maxFreqSum(s string) int {
	// Use slices as effectively as a hash map
	freqVowel := make([]int, 26)
	freqConsonant := make([]int, 26)

	// Iterate all characters
	for _, ch := range s {
		// Either add them to the vowel or consonant map
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			freqVowel[ch-'a']++
		} else {
			freqConsonant[ch-'a']++
		}
	}

	// Sum the maximum frequency of vowels and consonants
	return slices.Max(freqVowel) + slices.Max(freqConsonant)
}
