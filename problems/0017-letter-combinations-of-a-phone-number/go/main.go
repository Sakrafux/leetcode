package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations(""))   // []
	fmt.Println(letterCombinations("2"))  // ["a","b","c"]

}

// Solution
func letterCombinations(digits string) []string {
	// Handle the special empty case
	if len(digits) == 0 {
		return []string{}
	}

	// Build a hash-map of how to interpret a character
	digitToLetters := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	// In order to not handle the first iteration differently in the loop,
	// add an empty string
	res := make([][]rune, 1)
	res[0] = []rune{}

	for i, digit := range digits {
		letters := digitToLetters[digit]

		// For each iteration, discard the previous results
		newRes := make([][]rune, 0)
		// Iterate through the possible letters...
		for _, letter := range letters {
			// ...and append them to each currently known combination
			for _, item := range res {
				newItem := make([]rune, len(item)+1)
				copy(newItem, item)
				newItem[i] = letter
				newRes = append(newRes, newItem)
			}
		}
		res = newRes
	}

	resString := make([]string, len(res))
	for i, letters := range res {
		resString[i] = string(letters)
	}

	return resString
}
