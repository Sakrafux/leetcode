package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("([)]"))   // false
}

// Solution
func isValid(s string) bool {
	// Find the start for a ending bracket
	mapping := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// A stack of the opening brackets
	opens := make([]byte, 0)
	for i := range s {
		char := s[i]
		// If it is an opening bracket, push it onto the stack
		if char == '(' || char == '[' || char == '{' {
			opens = append(opens, char)
			// Otherwise, it must be a closing bracket
		} else {
			// Check if we have any openings at all
			lenOpens := len(opens)
			// If not, it is invalid
			if lenOpens == 0 {
				return false
			}
			// Check if the closing bracket fits the last openinig
			if opens[lenOpens-1] == mapping[char] {
				opens = opens[:lenOpens-1]
				// Otherwise it is invalid
			} else {
				return false
			}
		}
	}

	// At the end we must have emptied our stack
	return len(opens) == 0
}
