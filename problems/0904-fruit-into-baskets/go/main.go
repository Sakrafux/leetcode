package main

import (
	"fmt"
)

func main() {
	fmt.Println(totalFruit([]int{1, 2, 1}))          // 3
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))       // 3
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))    // 4
	fmt.Println(totalFruit([]int{1, 2, 3, 1, 2, 2})) // 3
}

// Solution
func totalFruit(fruits []int) int {
	// Keep a reference to the last time we found a fruit
	lastIndex := map[int]int{
		fruits[0]: 0,
	}

	// Due to the constraints there is at least 1 tree
	res := 1

	cur := 1
	// Remember the currently selected fruits up to 2
	selFruits := []int{fruits[0]}
	// Use a sliding window of the selected trees
	for l, r := 0, 1; r < len(fruits); r++ {
		newFruit := fruits[r]
		cur++
		// Update the index of the chosen fruit
		lastIndex[newFruit] = r

		// If we have less than 2 fruits, we don't have to consider anything
		if len(selFruits) < 2 {
			// Though only append it, if we don't have it yet
			if newFruit != selFruits[0] {
				selFruits = append(selFruits, newFruit)
			}
			// Otherwise, if it is a new fruit...
		} else if newFruit != selFruits[0] && newFruit != selFruits[1] {
			// ...we need to check which fruit we haven't seen for longer and set the start of the window behind it
			if li1, li2 := lastIndex[selFruits[0]], lastIndex[selFruits[1]]; li1 < li2 {
				l = li1 + 1
				selFruits[0] = newFruit
			} else {
				l = li2 + 1
				selFruits[1] = newFruit
			}
			// Also update the length of the current window
			cur = r - l + 1
		}

		// Keep only the longest one
		if cur > res {
			res = cur
		}
	}

	return res
}
