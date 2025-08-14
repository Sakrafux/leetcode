package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(minCost([]int{4, 2, 2, 2}, []int{1, 4, 1, 2}))                                                         // 1
	fmt.Println(minCost([]int{2, 3, 4, 1}, []int{3, 2, 5, 1}))                                                         // -1
	fmt.Println(minCost([]int{2, 4, 4, 1}, []int{3, 2, 3, 1}))                                                         // 2
	fmt.Println(minCost([]int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88}, []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8})) // 48
	fmt.Println(minCost([]int{4, 4, 4, 4, 3}, []int{5, 5, 5, 5, 3}))                                                   // 8
}

// Solution
func minCost(basket1 []int, basket2 []int) int64 {
	// A set to keep all fruit types
	elementsMap := make(map[int]bool)

	// Count the occurrences of each fruit in each basket
	freq1 := make(map[int]int)
	for _, num := range basket1 {
		freq := freq1[num]
		freq1[num] = freq + 1
		elementsMap[num] = true
	}
	freq2 := make(map[int]int)
	for _, num := range basket2 {
		freq := freq2[num]
		freq2[num] = freq + 1
		elementsMap[num] = true
	}

	minElement := math.MaxInt
	needSwap := make([]int, 0)
	// For each fruit, check if they need to be swapped
	for element := range elementsMap {
		value1 := freq1[element]
		value2 := freq2[element]

		// If the sum of a fruit type is odd, then we obviously can't split them equally
		if (value1+value2)%2 == 1 {
			return -1
		}
		// Otherwise, their difference is always even, which can be cleanly divided by 2,
		// which represents the amount of fruits we have to swap to achieve an even split
		if value1 != value2 {
			for range int(math.Abs(float64((value1 - value2) / 2))) {
				needSwap = append(needSwap, element)
			}
		}
		// Also keep in mind which element is the smallest
		if element < minElement {
			minElement = element
		}
	}

	// To ensure minimality, we need to sort the fruits to swap
	slices.Sort(needSwap)

	var res int64
	// The baskets must be equally unbalanced, i.e. every swap always has a partner
	// and because our cost calculation only cares about the fruit, we only need to traverse half the array
	for i := 0; i < len(needSwap)/2; i++ {
		// We can make use of the fact that we may use the smallest element available as a helper to swap 2 elements
		// in 2 operations instead of 1, which makes sense to do, if this is still cheaper than swapping the larget
		// elements directly
		res += int64(min(needSwap[i], 2*minElement))
	}

	return res
}
