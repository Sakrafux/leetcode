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
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2,
		&ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)) // [1,2,3,5]
	fmt.Println(removeNthFromEnd(&ListNode{1, nil}, 1))               // []
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, nil}}, 1)) // [1]
}

// Solution
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Save all the references into an array
	nodes := make([]*ListNode, 0)
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}

	// Handle the special case that the first node is removed
	sz := len(nodes)
	if n == sz {
		return head.Next
	}

	// Get the node before and after the one to remove
	prevNode := nodes[sz-n-1]
	var nextNode *ListNode = nil
	if n-1 > 0 {
		nextNode = nodes[sz-n+1]
	}

	// And remove the target node
	prevNode.Next = nextNode

	return head
}
