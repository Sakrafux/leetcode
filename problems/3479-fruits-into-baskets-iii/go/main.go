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
	// An array representation of a segment tree
	seg := make([]int, 4*n+1)
	// Recursively build the segment tree
	buildSegmentTree(seg, baskets, 1, 0, n-1)

	unplaced := 0

	// Iteratively try each fruit
	for _, fruit := range fruits {
		basketIndex := firstFittingBasket(seg, fruit, 1, 0, n-1)
		// If the index is n, we haven't found a fitting basket
		if basketIndex == n {
			unplaced++
			// Otherwise, ensure that the basket is marked
		} else {
			assign(seg, basketIndex, 0, 1, 0, n-1)
		}
	}

	return unplaced
}

// Set a parent node´s value to the maximum of its children (max segment tree)
func update(seg []int, index int) {
	seg[index] = max(seg[index*2], seg[index*2+1])
}

// Recursively build a segment tree of the baskets
// The index refers to the index in the segment tree array, while left and right correspond to the basket array
func buildSegmentTree(seg, baskets []int, index, left, right int) {
	// If it is a single element range, it is a leaf node
	if left == right {
		seg[index] = baskets[left]
		return
	}
	// Otherwise, we can split the range and continue building the tree recursively
	mid := (left + right) / 2
	buildSegmentTree(seg, baskets, index*2, left, mid)
	buildSegmentTree(seg, baskets, index*2+1, mid+1, right)
	update(seg, index)
}

// Assign a value to the corresponding leaf node and update the segment tree accordingly
// to ensure it´s still a max segment tree
func assign(seg []int, basketIndex, basketValue, segIndex, left, right int) {
	// Only update the correct leaf node
	if basketIndex < left || basketIndex > right {
		return
	}
	if left == right {
		seg[segIndex] = basketValue
		return
	}
	// Otherwise traverse the tree until you find it...
	mid := (left + right) / 2
	assign(seg, basketIndex, basketValue, segIndex*2, left, mid)
	assign(seg, basketIndex, basketValue, segIndex*2+1, mid+1, right)
	// ...and update accordingly
	update(seg, segIndex)
}

// Traverse the tree left-first until a fitting basket is found
func firstFittingBasket(seg []int, fruit, index, left, right int) int {
	// Since each parent is >= its children, we don't need to check the remaining tree, if it is too small
	if seg[index] < fruit {
		return right + 1
	}
	// Return the index of the correct leaf node
	if left == right {
		return left
	}
	mid := (left + right) / 2
	// Try left first
	leftBasketIndex := firstFittingBasket(seg, fruit, index*2, left, mid)
	if leftBasketIndex <= mid {
		return leftBasketIndex
	}
	// Only then try right
	return firstFittingBasket(seg, fruit, index*2+1, mid+1, right)
}
