package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}

// Solution
func lengthOfLongestSubstring(s string) int {
	charLastIndexMap := make(map[rune]int)

	maxLength := 0
	windowStartIndex := 0
	for i, char := range s {
		// We only care about a new character if we have already seen it...
		if charIndex, ok := charLastIndexMap[char]; ok {
			// ...and we only need to act upon it, if we have not already passed it
			windowStartIndex = int(math.Max(float64(windowStartIndex), float64(charIndex+1)))
		}
		charLastIndexMap[char] = i
		maxLength = int(math.Max(float64(maxLength), float64(i-windowStartIndex+1)))
	}

	return maxLength
}

func lengthOfLongestSubstring_hybrid(s string) int {
	letters := make(map[int]map[rune]bool)

	maxLength := 0
	for i, char := range s {
		letters[i] = make(map[rune]bool)
		for j, m := range letters {
			if _, ok := m[char]; ok {
				length := len(m)
				if length > maxLength {
					maxLength = length
				}
				delete(letters, j)
			} else {
				m[char] = true
			}
		}
	}

	for _, m := range letters {
		if length := len(m); length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

// This tries to create every possible substring and then check its validity and length
func lengthOfLongestSubstring_bruteForce(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	possibleSubstrings := make(map[string]int)

	for end := 1; end <= len(s); end++ {
		for start := 0; start < end; start++ {
			substring := s[start:end]
			possibleSubstrings[substring] = len(substring)
		}
	}

	maxLength := 0
	for substring, length := range possibleSubstrings {
		if checkSubstring(substring) && length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}

func checkSubstring(s string) bool {
	letters := make(map[rune]bool)
	for _, char := range s {
		if letters[char] {
			return false
		}
		letters[char] = true
	}
	return true
}
