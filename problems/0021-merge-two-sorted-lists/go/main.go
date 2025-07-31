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
	fmt.Println(mergeTwoLists(
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}})) // [1,1,2,3,4,4]
	fmt.Println(mergeTwoLists(nil, nil))               // []
	fmt.Println(mergeTwoLists(nil, &ListNode{0, nil})) // [0]
}

// Solution
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// A dummy node to handle the head uniformly
	dummy := &ListNode{}

	node := dummy
	// While both lists have nodes, check which is smaller and append that
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	// Check which list, if any, still has nodes and inherit them
	if list1 != nil {
		node.Next = list1
	} else if list2 != nil {
		node.Next = list2
	}

	// Skip the dummy start
	return dummy.Next
}
