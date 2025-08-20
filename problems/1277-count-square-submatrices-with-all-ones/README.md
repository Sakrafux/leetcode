# [1277. Count Square Submatrices with All Ones](https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/?envType=daily-question&envId=2025-08-20)

Given a <code>m * n</code> matrix of ones and zeros, return how many **square**  submatrices have all ones.

**Example 1:**

```
Input: matrix =
[
 [0,1,1,1],
 [1,1,1,1],
 [0,1,1,1]
]
Output: 15
Explanation: 
There are **10**  squares of side 1.
There are **4**  squares of side 2.
There is  **1**  square of side 3.
Total number of squares = 10 + 4 + 1 = **15** .
```

**Example 2:**

```
Input: matrix = 
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation: 
There are <b>6</b> squares of side 1.  
There is **1**  square of side 2. 
Total number of squares = 6 + 1 = <b>7</b>.
```

**Constraints:**

- <code>1 <= arr.length<= 300</code>
- <code>1 <= arr[0].length<= 300</code>
- <code>0 <= arr[i][j] <= 1</code>

## Solution

For this problem, we need to iterate through every cell once and built upon the knowledge with dynamic programming.
We simply analyze how many squares end at a particular cell. This is based on its neighbors by capturing the 
lowest order square among them, since this limits the highest order square available at the cell in question.

E.g. if every neighbor is the end of a 2x2 matrix, then the current cell is the end of a 3x3 matrix. If one of those
neighbors is only 1x1, then we can't end a 3x3 matrix, but only a 2x2 matrix.

By the end, every cell contains the information of how many squares end there, so we only need to sum. However, to
be most efficient, we can simply reuse the original matrix and since we never revisit a cell, we can sum up during
the iteration. Additionally, we don't need any additional check against the neighbors, since the minimum already 
implies a shape.

The resulting time complexity is `O(m*n)`, since we only iterate through the matrix once.