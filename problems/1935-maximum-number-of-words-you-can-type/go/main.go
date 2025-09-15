package main

import (
	"fmt"
)

func main() {
	fmt.Println(canBeTypedWords("hello world", "ad")) // 1
	fmt.Println(canBeTypedWords("leet code", "lt"))   // 1
	fmt.Println(canBeTypedWords("leet code", "e"))    // 0
}

// Solution
func canBeTypedWords(text string, brokenLetters string) int {
	// Map of the broken letters for fast checks
	blMap := make(map[rune]bool)
	for _, char := range brokenLetters {
		blMap[char] = true
	}

	words := 0
	wordSoFar := true
	// Iterate all characters of the string
	for _, char := range text {
		// If we reach a space, it is time to evaluate
		if char == ' ' {
			// If we were able to type the entire word, add one
			if wordSoFar {
				words++
			}
			// The next word starts out possible
			wordSoFar = true
			// Otherwise, we need to check whether we hit a broken letter
		} else if blMap[char] {
			// In that case we can't type the word
			wordSoFar = false
		}
	}
	// The last word is checked outside the loop
	if wordSoFar {
		words++
	}

	return words
}
