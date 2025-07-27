# [2. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/description/)

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** , and each of their nodes contains a single digit. Add the two numbers and return the sumas a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example 1:**
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg" style="width: 483px; height: 342px;">

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

**Constraints:**

- The number of nodes in each linked list is in the range <code>[1, 100]</code>.
- <code>0 <= Node.val <= 9</code>
- It is guaranteed that the list represents a number that does not have leading zeros.

## Solution

Since the result should also be a linked list, we don't need to bother with any conversions. Even more so, since the linked list starts at the least significant digit, which means we can simply move forward the list and correctly calculate the carry-over value.

The idea is now to simply advance both lists together and create a new node for our solution based on the sum, considering the carry-over value as well. Once at least one list has no next node, we check if one list still has nodes and traverse it similar to before until completion. At the end we simply need to check whether we still have a carry-over and, if yes, create a new node for it.

To keep the code simple without any special considerations for the first element, we initialize with a dummy node that then simply gets discarded at return.
