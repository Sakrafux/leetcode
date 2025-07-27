package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))          // 2
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))       // 2.5
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3, 4, 5, 6})) // 3.5
}

// Solution
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)

	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	n := n1 + n2
	left := (n1 + n2 + 1) / 2
	low, high := 0, n1

	for low <= high {
		mid1 := (low + high) / 2
		mid2 := left - mid1

		l1, l2, r1, r2 := math.MinInt, math.MinInt, math.MaxInt, math.MaxInt

		if mid1 < n1 {
			r1 = nums1[mid1]
		}
		if mid2 < n2 {
			r2 = nums2[mid2]
		}
		if (mid1 - 1) >= 0 {
			l1 = nums1[mid1-1]
		}
		if (mid2 - 1) >= 0 {
			l2 = nums2[mid2-1]
		}

		if l1 <= r2 && l2 <= r1 {
			if n%2 == 1 {
				return math.Max(float64(l1), float64(l2))
			} else {
				return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2
			}
		} else if l1 > r2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	panic("No solution")
}

func findMedianSortedArrays_2pointer(nums1 []int, nums2 []int) float64 {
	mergedArray := make([]int, 0, len(nums1)+len(nums2))
	for im, i1, i2 := 0, 0, 0; i1 < len(nums1) || i2 < len(nums2); im++ {
		if i1 == len(nums1) {
			mergedArray = append(mergedArray, nums2[i2])
			i2++
		} else if i2 == len(nums2) {
			mergedArray = append(mergedArray, nums1[i1])
			i1++
		} else {
			if nums1[i1] <= nums2[i2] {
				mergedArray = append(mergedArray, nums1[i1])
				i1++
			} else {
				mergedArray = append(mergedArray, nums2[i2])
				i2++
			}
		}
	}

	if len(mergedArray)%2 == 1 {
		return float64(mergedArray[len(mergedArray)/2])
	}
	i := len(mergedArray) / 2
	return float64(mergedArray[i-1]+mergedArray[i]) / 2
}
