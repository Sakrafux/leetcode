package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a"))            // false
	fmt.Println(isMatch("aa", "a*"))           // true
	fmt.Println(isMatch("ab", ".*"))           // true
	fmt.Println(isMatch("ab", ".*a"))          // false
	fmt.Println(isMatch("ab", ".*ab"))         // true
	fmt.Println(isMatch("abab", "ab.*ab"))     // true
	fmt.Println(isMatch("abaab", "ab.*ab"))    // true
	fmt.Println(isMatch("abaaab", "b*ab.*ab")) // true
}

// Solution
func isMatch(s string, p string) bool {
	// Initialize a dynamic programming table
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	// Empty string matches empty pattern
	dp[0][0] = true

	// Empty string matches any repeating * pattern
	for j := range p {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j-1]
		}
	}
	fmt.Println(dp[0])

	for i, sChar := range s {
		for j, pChar := range p {
			switch pChar {
			// For simple character matching, simply check the previous match
			case sChar, '.':
				dp[i+1][j+1] = dp[i][j]
			case '*':
				// Consider no occurences...
				dp[i+1][j+1] = dp[i+1][j-1]
				// ...or multiple occurences
				if rune(p[j-1]) == sChar || p[j-1] == '.' {
					dp[i+1][j+1] = dp[i+1][j+1] || dp[i][j+1]
				}
			}
		}
		fmt.Println(dp[i+1])
	}

	return dp[len(s)][len(p)]
}
