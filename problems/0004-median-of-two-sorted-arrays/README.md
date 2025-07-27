# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

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

This problem is interesting because the difficulty lies not in getting a solution but rather in the time complexity.

The target of `O(log(n+m))` already gives the hint that we shouldn't iterate linearly over the array(s), but rather multiply/divide the search space. The obvious contender to do so in arrays is binary search, which halves the search space each iteration.

However, the problem is that we don't have a coherent array we can comfortably use, but rather 2 seperate arrays and we can't merge them, because this would require linear time atleast. The idea is to treat both together as an imaginary total array, seperate this array into 2 equal partitions (or the left +1 in case of odd size) and keep pointers for both input arrays to mark the partitions.

We then need only to move the partition line pointer in one array using binary search; the other partition line pointer is forced to move correspondingly, because the partition sizes are clear from the beginning. For convenience's sake, we do the search always on the smaller array (which we always place first, swapping them if necessary).

The correctness of the partitions is determined by the fact that their surrounding elements, i.e. the first element left and right of the partition line, are in the correct order for both arrays. This means that the element right of the line in one array may not be smaller than the element left of the line in the other array. If this is the case, then the partitions are not correctly split between the arrays yet.

An example:

```
Start:
a1 = [1 3]
a2 = [2 4 5]
left = (2 + 3 + 1) / 2 = 3
low = 0
high = 2

Iteration 1:
mid1 = 1 -> mid2 = 2
a1 = [1 | 3]
a2 = [2 4 | 5]

l1 <= r2 <=> 1 <= 5
l2 <= r1 <=> 4 <= 3 // not true

-> low = 2

Iteration 2:
mid1 = 2 -> mid2 = 1
a1 = [1 3 |]
a2 = [2 | 4 5]

l1 <= r2 <=> 3 <= 4
l2 <= r1 <=> 2 <= Inf

Median = max(l1, l2) = max(3, 2) = 3
```

Another example:

```
Start:
a1 = [4 7 8]
a2 = [1 2 3 5 6]
left = (3 + 5 + 1) / 2 = 4
low = 0
high = 3

Iteration 1:
mid1 = 1 -> mid2 = 3
a1 = [4 | 6 7]
a2 = [1 2 3 | 3 6]

l1 <= r2 <=> 4 <= 3 // not true
l2 <= r1 <=> 3 <= 6

-> high = 0

Iteration 2:
mid1 = 0 -> mid2 = 4
a1 = [| 4 6 7]
a2 = [1 2 3 3 | 6]

l1 <= r2 <=> -Inf <= 6
l2 <= r1 <=> 3 <= 4

Median = (max(l1, l2) + min(r1, r2)) / 2 = (max(-Inf, 3) + min(6, 4)) / 2 = 3.5
```

### Another approach

A much easier approach is to simply merge the 2 arrays using 2 pointers that move according to who points to the smaller element. This takes `O(n+m)` and the merged array can then be simplied used to calculate the median.
