package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}

// Solution
func longestCommonPrefix(strs []string) string {
	res := make([]byte, 0)

	for i := 0; ; i++ {
		// Don't continue if first string is already finished
		if len(strs[0]) == i {
			break
		}
		// Get the character all have to have from the first string
		char := strs[0][i]
		// Iterate all strings...
		for _, s := range strs {
			// ...and check whether they are already finished or have a different character
			if len(s) == i || s[i] != char {
				return string(res)
			}
		}
		// If everything cleared, append the character
		res = append(res, char)
	}

	return string(res)
}
