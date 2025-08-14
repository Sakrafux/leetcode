package main

import (
	"fmt"
)

func main() {
	fmt.Println(numOfUnplacedFruits([]int{4, 2, 5}, []int{3, 5, 4})) // 1
	fmt.Println(numOfUnplacedFruits([]int{3, 6, 1}, []int{6, 4, 7})) // 0
}

// Solution
func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(baskets)
	seg := make([]int, 4*n+1)
	buildSegmentTree(seg, baskets, 1, 0, n-1)

	unplaced := 0

	for _, fruit := range fruits {
		basketIndex := firstFittingBasket(seg, fruit, 1, 0, n-1)
		if basketIndex == n {
			unplaced++
		} else {
			assign(seg, basketIndex, 0, 1, 0, n-1)
		}
	}

	return unplaced
}

func update(seg []int, index int) {
	seg[index] = max(seg[index*2], seg[index*2+1])
}

func buildSegmentTree(seg []int, baskets []int, index int, left int, right int) {
	if left == right {
		seg[index] = baskets[left]
		return
	}
	mid := (left + right) / 2
	buildSegmentTree(seg, baskets, index*2, left, mid)
	buildSegmentTree(seg, baskets, index*2+1, mid+1, right)
	update(seg, index)
}

func assign(seg []int, basketIndex int, basketValue int, segIndex int, left int, right int) {
	if basketIndex < left || basketIndex > right {
		return
	}
	if left == right {
		seg[segIndex] = basketValue
		return
	}
	mid := (left + right) / 2
	assign(seg, basketIndex, basketValue, segIndex*2, left, mid)
	assign(seg, basketIndex, basketValue, segIndex*2+1, mid+1, right)
	update(seg, segIndex)
}

func firstFittingBasket(seg []int, fruit int, index int, left int, right int) int {
	if seg[index] < fruit {
		return right + 1
	}
	if left == right {
		return left
	}
	mid := (left + right) / 2
	leftBasketIndex := firstFittingBasket(seg, fruit, index*2, left, mid)
	if leftBasketIndex <= mid {
		return leftBasketIndex
	}
	return firstFittingBasket(seg, fruit, index*2+1, mid+1, right)
}
