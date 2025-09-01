package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxAverageRatio([][]int{{1, 2}, {3, 5}, {2, 2}}, 2))                                                // 0.78333
	fmt.Println(maxAverageRatio([][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}, 4))                                       // 0.53485
	fmt.Println(maxAverageRatio([][]int{{583, 868}, {783, 822}, {65, 262}, {121, 508}, {461, 780}, {484, 668}}, 8)) // 0.57472
}

// Solution
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	// Create a 1-indexed heap
	heap := make([][]int, len(classes)+1)

	// Fill the heap up
	for i := 1; i <= len(classes); i++ {
		heap[i] = classes[i-1]
		heapifyUp(heap, i)
	}

	// Add an extra student to the maximum element and then heapify it down
	for range extraStudents {
		heap[1][0] += 1
		heap[1][1] += 1
		heapifyDown(heap, 1)
	}

	// Calculate the sum of all ratios for the average
	sumRatio := 0.0
	for i := 1; i <= len(classes); i++ {
		item := heap[i]
		sumRatio += float64(item[0]) / float64(item[1])
	}

	// Calculate the average ratio based on the sum of ratios
	return sumRatio / float64(len(classes))
}

// Helper function to calculate the increase in ratio by adding a student
func ratioIncrease(class []int) float64 {
	return (float64(class[0]+1) / float64(class[1]+1)) - (float64(class[0]) / float64(class[1]))
}

// Transport an element up the heap if it has a larger ratio increase than its parent
func heapifyUp(heap [][]int, i int) {
	if i == 1 {
		return
	}

	parent := heap[i/2]
	child := heap[i]

	// Check if the parent has a smaller ratio than the child
	if ratioIncrease(parent) < ratioIncrease(child) {
		heap[i/2], heap[i] = child, parent
		heapifyUp(heap, i/2)
	}
}

// Transport an element down the heap if it has a smaller ratio increase, switching it with the larger child
func heapifyDown(heap [][]int, i int) {
	// Check for the first child
	if i*2 >= len(heap) {
		return
	}

	parent := heap[i]
	child1 := heap[i*2]

	// Check if the second child exists
	if i*2+1 < len(heap) {
		child2 := heap[i*2+1]
		// Check if the second child has the larger ratio and compare it with the parent
		if ratioIncrease(child1) < ratioIncrease(child2) {
			if ratioIncrease(parent) < ratioIncrease(child2) {
				heap[i*2+1], heap[i] = parent, child2
				heapifyDown(heap, i*2+1)
			}
			return
		}
	}

	// Check if the first child is larger
	if ratioIncrease(parent) < ratioIncrease(child1) {
		heap[i*2], heap[i] = parent, child1
		heapifyDown(heap, i*2)
	}
}
