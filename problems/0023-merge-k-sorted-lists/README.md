# [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/description/)

You are given an array of <code>k</code> linked-lists <code>lists</code>, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

**Example 1:**

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted linked list:
1->1->2->3->4->4->5->6
```

**Example 2:**

```
Input: lists = []
Output: []
```

**Example 3:**

```
Input: lists = [[]]
Output: []
```

**Constraints:**

- <code>k == lists.length</code>
- <code>0 <= k <= 10^4</code>
- <code>0 <= lists[i].length <= 500</code>
- <code>-10^4 <= lists[i][j] <= 10^4</code>
- <code>lists[i]</code> is sorted in **ascending order** .
- The sum of <code>lists[i].length</code> will not exceed <code>10^4</code>.

## Solution

The solution to the problem is min-heap, i.e. a binary tree, where the parent is smaller than its children. The top element is the smallest element in the heap. The idea is that if repeatedly remove the top, we naturally get a sorted list.

An array can be used as a heap, where for each node with index `i`, its children are at index `2i + 1` and `2i + 2`, and its parent is at `(i-1) / 2`.

For inserting a node into a heap, we append it at the end and then heapify it up, which means we repeatedly check it against its parent and swap them, if it is smaller, or stop otherwise.

Similarly, when removing the head we move the last node to the top and then heapify it down, which means repeatedly checking it against its children and swapping them, if it is larger, or stop otherwise.

The time complexity of this is `O(n*k*log(n*k))`, where `n` is the maximum length of any list and `k` the number of lists. Since heap insert/removal has `O(size)` and `size` is the maximum number of elements given by `n*k`.
