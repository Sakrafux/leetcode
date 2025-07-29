package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))                                        // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 1, 1}))                                                    // []
	fmt.Println(threeSum([]int{0, 0, 0}))                                                    // [[0,0,0]]
	fmt.Println(threeSum([]int{0, 0, 0, 0, 0, 0}))                                           // [[0,0,0]]
	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10})) // [[-10,5,5],[-5,0,5],[-4,2,2],[-3,-2,5],[-3,1,2],[-2,0,2]]
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))                       // [[-4,0,4],[-4,1,3],[-3,-1,4],[-3,0,3],[-3,1,2],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]
}

// Solution
func threeSum(nums []int) [][]int {
	// Sort the array for assurance of ordinality
	slices.Sort(nums)

	res := make([][]int, 0)

	// Fix index i per iteration
	for i, num1 := range nums {
		// If we already saw the combination, continue
		if i > 0 && nums[i-1] == num1 {
			continue
		}

		// 2 pointers at both ends of the remaining window
		// that move until they overtake each other
		for j, k := i+1, len(nums)-1; j < k; {
			num2, num3 := nums[j], nums[k]
			sum := num1 + num2 + num3

			// If the sum is smaller than 0, we need to increase it,
			// which can be achieved by pushing the left pointer
			// towards larger numbers (i.e. rightwards)
			if sum < 0 {
				j++
				// If the sum is bigger than 0, the same principal applies
				// to the right pointer
			} else if sum > 0 {
				k--
				// If it is 0, we found a solution
			} else {
				res = append(res, []int{num1, num2, num3})
				// Try the next combination
				j++

				// Ensure that we get a new combination
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	}

	return res
}

func _threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	usedComb := make(map[int]map[int]bool)

	for i, num1 := range nums {
		lastIndex := make(map[int]int)
		for j := len(nums) - 1; j > i; j-- {
			num2 := nums[j]
			value := 0 - num1 - num2
			if k, ok := lastIndex[value]; ok {
				arr := sorted3Array(num1, num2, nums[k])
				isAppend := true

				if usedComb2, ok := usedComb[arr[0]]; ok {
					if _, ok := usedComb2[arr[1]]; ok {
						isAppend = false
					} else {
						usedComb2[arr[1]] = true
					}
				} else {
					usedComb[arr[0]] = map[int]bool{arr[1]: true}
				}

				if isAppend {
					res = append(res, arr)
				}
			}
			lastIndex[num2] = j
		}
	}

	return res
}

func sorted3Array(a, b, c int) []int {
	if a < b && a < c {
		if b < c {
			return []int{a, b, c}
		} else {
			return []int{a, c, b}
		}
	} else {
		if a < b {
			return []int{c, a, b}
		} else if a < c {
			return []int{b, a, c}
		} else if b < c {
			return []int{b, c, a}
		} else {
			return []int{c, b, a}
		}
	}
}
