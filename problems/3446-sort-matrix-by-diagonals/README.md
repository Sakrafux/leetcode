# [3446. Sort Matrix by Diagonals](https://leetcode.com/problems/sort-matrix-by-diagonals/description/?envType=daily-question&envId=2025-08-28)

You are given an <code>n x n</code> square matrix of integers <code>grid</code>. Return the matrix such that:

- The diagonals in the **bottom-left triangle**  (including the middle diagonal) are sorted in **non-increasing order** .
- The diagonals in the **top-right triangle**  are sorted in **non-decreasing order** .

**Example 1:**

<div class="example-block">
Input: grid = [[1,7,3],[9,8,2],[4,5,6]]

Output: [[8,2,3],[9,6,7],[4,5,1]]

Explanation:

<img alt="" src="https://assets.leetcode.com/uploads/2024/12/29/4052example1drawio.png" style="width: 461px; height: 181px;">

The diagonals with a black arrow (bottom-left triangle) should be sorted in non-increasing order:

- <code>[1, 8, 6]</code> becomes <code>[8, 6, 1]</code>.
- <code>[9, 5]</code> and <code>[4]</code> remain unchanged.

The diagonals with a blue arrow (top-right triangle) should be sorted in non-decreasing order:

- <code>[7, 2]</code> becomes <code>[2, 7]</code>.
- <code>[3]</code> remains unchanged.

**Example 2:**

<div class="example-block">
Input: grid = [[0,1],[1,2]]

Output: [[2,1],[1,0]]

Explanation:

<img alt="" src="https://assets.leetcode.com/uploads/2024/12/29/4052example2adrawio.png" style="width: 383px; height: 141px;">

The diagonals with a black arrow must be non-increasing, so <code>[0, 2]</code> is changed to <code>[2, 0]</code>. The other diagonals are already in the correct order.

**Example 3:**

<div class="example-block">
Input: grid = [[1]]

Output: [[1]]

Explanation:

Diagonals with exactly one element are already in order, so no changes are needed.

**Constraints:**

- <code>grid.length == grid[i].length == n</code>
- <code>1 <= n <= 10</code>
- <code>-10^5 <= grid[i][j] <= 10^5</code>

## Solution

For this problem, it is important to consider that we do not care about single cells but only entire diagonals. 
Furthermore, each diagonal stands by itself and is not impacted by other diagonals.
This means, we simply iterate `n`-times for all the lower diagonals, for each iteration extracting the diagonal into an array, 
sorting it, and then writing the sorted array back into the matrix. The upper diagonals work similarly, only for slightly
shifted positions and sorting.

The time complexity is `O(n²*log(n))`, because on a baseline we need to iterate all diagonals, which is `O(2n)` i.e. `O(n)`.
Inside each iteration we then iterate over a diagonal in `O(n)` (two times) and sort it in `O(n*log(n))`.