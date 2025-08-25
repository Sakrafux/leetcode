# [498. Diagonal Traverse](https://leetcode.com/problems/diagonal-traverse/description/?envType=daily-question&envId=2025-08-25)

Given an <code>m x n</code> matrix <code>mat</code>, return an array of all the elements of the array in a diagonal order.

**Example 1:**
<img alt="" src="https://assets.leetcode.com/uploads/2021/04/10/diag1-grid.jpg" style="width: 334px; height: 334px;">

```
Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]
```

**Example 2:**

```
Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]
```

**Constraints:**

- <code>m == mat.length</code>
- <code>n == mat[i].length</code>
- <code>1 <= m, n <= 10^4</code>
- <code>1 <= m * n <= 10^4</code>
- <code>-10^5 <= mat[i][j] <= 10^5</code>

## Solution

It is rather simple to move diagonally through a matrix, with the only issue being the handling of the edges, while 
remembering the direction of travel. In either direction, there are 3 edge cases: only the horizontal edge, only the
vertical edge, or the corner. Those need to be handled and subsequently change the direction of travel.

As we are traversing the entire matrix only once, we have a time complexity of `O(m*n)`.