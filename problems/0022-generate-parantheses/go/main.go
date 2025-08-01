package main

import (
	"fmt"
)

func main() {
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(generateParenthesis(1)) // ["()"]
	fmt.Println(generateParenthesis(4)) // ["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
}

// Solution
func generateParenthesis(n int) []string {
	// Keep all the patterns generated in a map by the number of parantheses
	patternsByNumber := map[int][]string{
		// Start with the initial case
		1: {"()"},
	}

	// Generate a new entry for the map
	for i := 2; i <= n; i++ {
		// An array for the results for the current number
		curPatterns := make([]string, 0)
		// All simple wrappings of the previous results are part of the solution
		for _, pattern := range patternsByNumber[i-1] {
			curPatterns = append(curPatterns, "("+pattern+")")
		}

		// Use a hash map to ensure uniqueness of the solutions
		uniquePatterns := make(map[string]bool)
		// Try all the combinations of previous patterns that have the correct number of parantheses
		for j := 1; j < i; j++ {
			// Get a layer...
			patterns1 := patternsByNumber[j]
			// ...and its necessary counterpart
			patterns2 := patternsByNumber[i-j]

			// Try all the combinations, as they are necessarily valid
			for _, pattern1 := range patterns1 {
				for _, pattern2 := range patterns2 {
					uniquePatterns[pattern1+pattern2] = true
					uniquePatterns[pattern2+pattern1] = true
				}
			}
		}

		// Append the keys of the uniqueness map
		for pattern := range uniquePatterns {
			curPatterns = append(curPatterns, pattern)
		}
		// Save the current number's solution
		patternsByNumber[i] = curPatterns
	}

	// Take the solution from our pattern map
	return patternsByNumber[n]
}
