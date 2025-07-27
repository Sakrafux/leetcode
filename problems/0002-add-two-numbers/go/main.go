package main

import "fmt"

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
	fmt.Println(addTwoNumbers(&ListNode{2, &ListNode{4,
		&ListNode{3, nil}}}, &ListNode{5, &ListNode{6,
		&ListNode{4, nil}}})) // [7 0 8]
	fmt.Println(addTwoNumbers(&ListNode{0, nil}, &ListNode{0, nil})) // [0]
	fmt.Println(addTwoNumbers(&ListNode{9, &ListNode{9,
		&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9,
			&ListNode{9, nil}}}}}}}, &ListNode{9, &ListNode{9,
		&ListNode{9, &ListNode{9, nil}}}})) // [8 9 9 9 0 0 0 1]
}

// Solution
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	n := res
	rest := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		n.Next = &ListNode{}
		n = n.Next
		n.Val, rest = sumWithRest(l1.Val, l2.Val, rest)
	}

	var l *ListNode
	if l1 != nil {
		l = l1
	} else if l2 != nil {
		l = l2
	}

	for ; l != nil; l = l.Next {
		n.Next = &ListNode{}
		n = n.Next
		n.Val, rest = sumWithRest(l.Val, 0, rest)
	}

	if rest > 0 {
		n.Next = &ListNode{rest, nil}
	}

	return res.Next
}

func sumWithRest(i1, i2, rest int) (int, int) {
	sum := i1 + i2 + rest
	if sum > 9 {
		return sum - 10, 1
	}
	return sum, 0
}
