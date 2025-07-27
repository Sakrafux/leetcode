package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babad")) // bab
	fmt.Println(longestPalindrome("cbbd"))  // bb
	fmt.Println(longestPalindrome("ccc"))   // ccc
}

// Solution
func longestPalindrome(s string) string {
	res := string(s[0])
	maxLength := 1

	// Iterate over each character as a potential center
	for i := 0; i < len(s)-1; i++ {
		// Try steps from the center until we hit a boundary on either side
		for step := 0; i-step >= 0 && i+step < len(s); step++ {
			// Are the new characters the same, i.e. is it (still) a palindrome?
			if s[i-step] == s[i+step] {
				// If yes, compare to previous best
				if windowSize := step*2 + 1; windowSize > maxLength {
					maxLength = windowSize
					res = s[i-step : i+step+1]
				}
				// If not, don't bother going more steps
			} else {
				break
			}
		}

		// The same needs to be done for 0/2 character centers
		// Notice the i+1 for the right boundary at each usage
		for step := 0; i-step >= 0 && i+1+step < len(s); step++ {
			if s[i-step] == s[i+1+step] {
				if windowSize := step*2 + 2; windowSize > maxLength {
					maxLength = windowSize
					res = s[i-step : i+step+2]
				}
			} else {
				break
			}
		}
	}

	return res
}
