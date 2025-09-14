package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spellchecker([]string{"KiTe", "kite", "hare", "Hare"}, []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"})) // ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]
	fmt.Println(spellchecker([]string{"yellow"}, []string{"YellOw"}))                                                                                             // ["yellow"]
}

// Solution
func spellchecker(wordlist []string, queries []string) []string {
	// Create maps for fast lookup of wordlist words
	wordMap := make(map[string]bool)
	// lower-case versions mapping to the first occurrence in the wordlist
	wordLowerMap := make(map[string]string)
	// no-vowel versions mapping to the first occurrence in the wordlist
	wordNoVowelMap := make(map[string]string)
	for _, word := range wordlist {
		wordLower := strings.ToLower(word)
		noVowelBytes := []byte(wordLower)
		// Replace all vowels with spaces
		for j, char := range noVowelBytes {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				noVowelBytes[j] = ' '
			}
		}
		noVowelWord := string(noVowelBytes)
		if _, ok := wordMap[word]; !ok {
			wordMap[word] = true
		}
		if _, ok := wordLowerMap[wordLower]; !ok {
			wordLowerMap[wordLower] = word
		}
		if _, ok := wordNoVowelMap[noVowelWord]; !ok {
			wordNoVowelMap[noVowelWord] = word
		}
	}

	output := make([]string, len(queries))

	// Iterate all query strings
	for i, query := range queries {
		// Check for exact match
		if wordMap[query] {
			output[i] = query
			continue
		}

		// Check for capitalization match
		queryLower := strings.ToLower(query)
		if word, ok := wordLowerMap[queryLower]; ok {
			output[i] = word
			continue
		}

		// Check for vowel match
		noVowelWord := []byte(queryLower)
		for j, char := range queryLower {
			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				noVowelWord[j] = ' '
			}
		}
		queryNoVowel := string(noVowelWord)
		if word, ok := wordNoVowelMap[queryNoVowel]; ok {
			output[i] = word
			continue
		}

		// Otherwise no match
		output[i] = ""
	}

	return output
}
