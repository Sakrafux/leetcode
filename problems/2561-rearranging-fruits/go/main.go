package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(minCost([]int{4, 2, 2, 2}, []int{1, 4, 1, 2}))                                                         // 1
	fmt.Println(minCost([]int{2, 3, 4, 1}, []int{3, 2, 5, 1}))                                                         // -1
	fmt.Println(minCost([]int{84, 80, 43, 8, 80, 88, 43, 14, 100, 88}, []int{32, 32, 42, 68, 68, 100, 42, 84, 14, 8})) // 48
	fmt.Println(minCost([]int{4, 4, 4, 4, 3}, []int{5, 5, 5, 5, 3}))                                                   // 8
}

// Solution
func minCost(basket1 []int, basket2 []int) int64 {
	elementsMap := make(map[int]bool)

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
	for element := range elementsMap {
		value1 := freq1[element]
		value2 := freq2[element]

		if (value1+value2)%2 == 1 {
			return -1
		}
		if value1 != value2 {
			for range int(math.Abs(float64((value1 - value2) / 2))) {
				needSwap = append(needSwap, element)
			}
		}
		if element < minElement {
			minElement = element
		}
	}

	slices.Sort(needSwap)

	var res int64
	for i := 0; i < len(needSwap)/2; i++ {
		res += int64(math.Min(float64(needSwap[i]), float64(2*minElement)))
	}

	return res
}
