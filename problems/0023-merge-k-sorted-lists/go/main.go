package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	if node.Next == nil {
		return fmt.Sprint(node.Val)
	}
	return fmt.Sprintf("%v %v", node.Val, node.Next)
}

func main() {
	fmt.Println(mergeKLists([]*ListNode{
		{1, &ListNode{4, &ListNode{5, nil}}},
		{1, &ListNode{3, &ListNode{4, nil}}},
		{2, &ListNode{6, nil}},
	})) // [1,1,2,3,4,4,5,6]
	fmt.Println(mergeKLists([]*ListNode{}))    // []
	fmt.Println(mergeKLists([]*ListNode{nil})) // []
	fmt.Println(mergeKLists([]*ListNode{{-2, &ListNode{-1, &ListNode{-1, &ListNode{-1, nil}}}}, nil}))
}

// Solution
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	heap := make([]*ListNode, 0)

	// Add all elements to the heap
	for _, list := range lists {
		for node := list; node != nil; node = node.Next {
			heap = insertNode(heap, node)
		}
	}

	// An empty heap means no data therefore return nil
	if len(heap) == 0 {
		return nil
	}

	res := &ListNode{}
	list := res

	// Iteratively remove from the top of the heap until its empty
	// This sorts the result automatically
	var node *ListNode
	node, heap = removeHead(heap)
	for node != nil {
		list.Next = node
		list = list.Next
		node, heap = removeHead(heap)
	}

	return res.Next
}

func insertNode(heap []*ListNode, node *ListNode) []*ListNode {
	// Append the node to the end of heap...
	heap = append(heap, node)

	// ...and then heapify it up
	for i := len(heap) - 1; i > 0; {
		// The parent node's index
		p := (i - 1) / 2

		// If the node is smaller than its parent, swap them
		if heap[i].Val < heap[p].Val {
			heap[i], heap[p] = heap[p], heap[i]
			i = p
			// Otherwise, it is at the right place
		} else {
			break
		}
	}

	return heap
}

func removeHead(heap []*ListNode) (*ListNode, []*ListNode) {
	n := len(heap) - 1

	// Return nil, if the heap is empty
	if n < 0 {
		return nil, heap
	}

	// Get the top of the heap...
	res := heap[0]
	// ...and ensure its clean, i.e. forgets its previous successor
	res.Next = nil

	// Move the last element to the top of the heap...
	heap[0] = heap[n]

	// ...and heapify it down
	for i := 0; i < n; {
		val := heap[i].Val

		c1 := i*2 + 1
		c2 := i*2 + 2

		// Check if the index of the first child is within bounds
		if c1 < n {
			// Check if the index of the second child is within bounds
			if c2 < n {
				// Check if the first child is smaller
				if heap[c1].Val < heap[c2].Val {
					// Then check if we need to swap
					if heap[c1].Val < val {
						heap[i], heap[c1] = heap[c1], heap[i]
						i = c1
						// If not, then we are at the right place
					} else {
						break
					}
				} else {
					// Otherwise, check if we need to swap with the second child...
					if heap[c2].Val < val {
						heap[i], heap[c2] = heap[c2], heap[i]
						i = c2
						// ...or are at the right place
					} else {
						break
					}
				}
			} else {
				// Otherwise, there is only a single swap possible
				if heap[c1].Val < val {
					heap[i], heap[c1] = heap[c1], heap[i]
				}
				break
			}
			// Otherwise, it's already as low as possible
		} else {
			break
		}
	}

	// Return the top node and the now-smaller heap
	return res, heap[:n]
}
