package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(minimumTeachings(2, [][]int{{1}, {2}, {1, 2}}, [][]int{{1, 2}, {1, 3}, {2, 3}}))                 // 1
	fmt.Println(minimumTeachings(3, [][]int{{2}, {1, 3}, {1, 2}, {3}}, [][]int{{1, 4}, {1, 2}, {3, 4}, {2, 3}})) // 2
	fmt.Println(minimumTeachings(11,
		[][]int{{3, 11, 5, 10, 1, 4, 9, 7, 2, 8, 6}, {5, 10, 6, 4, 8, 7}, {6, 11, 7, 9}, {11, 10, 4}, {6, 2, 8, 4, 3}, {9, 2, 8, 4, 6, 1, 5, 7, 3, 10}, {7, 5, 11, 1, 3, 4}, {3, 4, 11, 10, 6, 2, 1, 7, 5, 8, 9}, {8, 6, 10, 2, 3, 1, 11, 5}, {5, 11, 6, 4, 2}},
		[][]int{{7, 9}, {3, 7}, {3, 4}, {2, 9}, {1, 8}, {5, 9}, {8, 9}, {6, 9}, {3, 5}, {4, 5}, {4, 9}, {3, 6}, {1, 7}, {1, 3}, {2, 8}, {2, 6}, {5, 7}, {4, 6}, {5, 8}, {5, 6}, {2, 7}, {4, 8}, {3, 8}, {6, 8}, {2, 5}, {1, 4}, {1, 9}, {1, 6}, {6, 7}},
	)) // 0
	fmt.Println(minimumTeachings(17,
		[][]int{{4, 7, 2, 14, 6}, {15, 13, 6, 3, 2, 7, 10, 8, 12, 4, 9}, {16}, {10}, {10, 3}, {4, 12, 8, 1, 16, 5, 15, 17, 13}, {4, 13, 15, 8, 17, 3, 6, 14, 5, 10}, {11, 4, 13, 8, 3, 14, 5, 7, 15, 6, 9, 17, 2, 16, 12}, {4, 14, 6}, {16, 17, 9, 3, 11, 14, 10, 12, 1, 8, 13, 4, 5, 6}, {14}, {7, 14}, {17, 15, 10, 3, 2, 12, 16, 14, 1, 7, 9, 6, 4}},
		[][]int{{4, 11}, {3, 5}, {7, 10}, {10, 12}, {5, 7}, {4, 5}, {3, 8}, {1, 5}, {1, 6}, {7, 8}, {4, 12}, {2, 4}, {8, 9}, {3, 10}, {4, 7}, {5, 12}, {4, 9}, {1, 4}, {2, 8}, {1, 2}, {3, 4}, {5, 10}, {2, 7}, {1, 7}, {1, 8}, {8, 10}, {1, 9}, {1, 10}, {6, 7}, {3, 7}, {8, 12}, {7, 9}, {9, 11}, {2, 5}, {2, 3}},
	)) // 4
}

// Solution
func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	// Create a map of all the languages a user speaks (due to the incremental integer keys, we can use an array instead)
	userLanguages := make([][]bool, len(languages))
	for i, l := range languages {
		languageMap := make([]bool, n)
		for _, j := range l {
			languageMap[j-1] = true
		}
		userLanguages[i] = languageMap
	}

	// Create a map of all the friendships that don't need a new language
	commonLanguageFriends := make([]bool, len(friendships))
	countCLF := 0
	// For each friendship...
	for f, friendship := range friendships {
		// ...try each language
		for i := range n {
			// Do both users understand the language?
			if userLanguages[friendship[0]-1][i] == true && userLanguages[friendship[1]-1][i] == true {
				commonLanguageFriends[f] = true
				countCLF++
				break
			}
		}
	}
	// If this concerns all friendships, no user is taught a language
	if countCLF == len(friendships) {
		return 0
	}

	// Keep track of the number of users that need to be taught by language
	usersPerLanguage := make([]int, n)

	// Iterate each language
	for i := range n {
		// Then iterate all friendships
		for f, friendship := range friendships {
			// If they speak the same language, skip it
			if commonLanguageFriends[f] == true {
				continue
			}

			u1l := userLanguages[friendship[0]-1]
			u2l := userLanguages[friendship[1]-1]
			// If any of them don't speak a language, they would need to be taught this language
			if u1l[i] == false || u2l[i] == false {
				// Either both or only one user needs to be taught
				if u1l[i] == u2l[i] {
					usersPerLanguage[i] += 2
				} else {
					usersPerLanguage[i] += 1
				}
				// Remember that they now speak the language and only need to be taught once
				u1l[i] = true
				u2l[i] = true
			}
		}
	}

	// Return the language that needs to be taught the least
	return slices.Min(usersPerLanguage)
}
