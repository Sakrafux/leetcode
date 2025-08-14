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
	lastIndex := map[int]int{
		fruits[0]: 0,
	}

	res := 1

	cur := 1
	selFruits := []int{fruits[0]}
	for l, r := 0, 1; r < len(fruits); r++ {
		newFruit := fruits[r]
		cur++
		lastIndex[newFruit] = r

		if len(selFruits) < 2 {
			if newFruit != selFruits[0] {
				selFruits = append(selFruits, newFruit)
			}
		} else if newFruit != selFruits[0] && newFruit != selFruits[1] {
			if li1, li2 := lastIndex[selFruits[0]], lastIndex[selFruits[1]]; li1 < li2 {
				l = li1 + 1
				selFruits[0] = newFruit
			} else {
				l = li2 + 1
				selFruits[1] = newFruit
			}
			cur = r - l + 1
		}

		if cur > res {
			res = cur
		}
	}

	return res
}
