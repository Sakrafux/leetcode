# [3459. Length of Longest V-Shaped Diagonal Segment](https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/description/?envType=daily-question&envId=2025-08-27)

You are given a 2D integer matrix <code>grid</code> of size <code>n x m</code>, where each element is either <code>0</code>, <code>1</code>, or <code>2</code>.

A **V-shaped diagonal segment**  is defined as:

- The segment starts with <code>1</code>.
- The subsequent elements follow this infinite sequence: <code>2, 0, 2, 0, ...</code>.
- The segment:

- Starts **along**  a diagonal direction (top-left to bottom-right, bottom-right to top-left, top-right to bottom-left, or bottom-left to top-right).
- Continues the** sequence**  in the same diagonal direction.
- Makes** at most one clockwise 90-degree** ** turn**  to another diagonal direction while **maintaining**  the sequence.

<img alt="" src="https://assets.leetcode.com/uploads/2025/01/11/length_of_longest3.jpg" style="width: 481px; height: 202px;">

Return the **length**  of the **longest**  **V-shaped diagonal segment** . If no valid segment exists, return 0.

**Example 1:**

<div class="example-block">
Input: grid = [[2,2,1,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]

Output: 5

Explanation:

<img alt="" src="https://assets.leetcode.com/uploads/2024/12/09/matrix_1-2.jpg" style="width: 201px; height: 192px;">

The longest V-shaped diagonal segment has a length of 5 and follows these coordinates: <code>(0,2) → (1,3) → (2,4)</code>, takes a **90-degree clockwise turn**  at <code>(2,4)</code>, and continues as <code>(3,3) → (4,2)</code>.

**Example 2:**

<div class="example-block">
Input: grid = [[2,2,2,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]

Output: 4

Explanation:

**<img alt="" src="https://assets.leetcode.com/uploads/2024/12/09/matrix_2.jpg" style="width: 201px; height: 201px;">**

The longest V-shaped diagonal segment has a length of 4 and follows these coordinates: <code>(2,3) → (3,2)</code>, takes a **90-degree clockwise turn**  at <code>(3,2)</code>, and continues as <code>(2,1) → (1,0)</code>.

**Example 3:**

<div class="example-block">
Input: grid = [[1,2,2,2,2],[2,2,2,2,0],[2,0,0,0,0],[0,0,2,2,2],[2,0,0,2,0]]

Output: 5

Explanation:

**<img alt="" src="https://assets.leetcode.com/uploads/2024/12/09/matrix_3.jpg" style="width: 201px; height: 201px;">**

The longest V-shaped diagonal segment has a length of 5 and follows these coordinates: <code>(0,0) → (1,1) → (2,2) → (3,3) → (4,4)</code>.

**Example 4:**

<div class="example-block">
Input: grid = [[1]]

Output: 1

Explanation:

The longest V-shaped diagonal segment has a length of 1 and follows these coordinates: <code>(0,0)</code>.

**Constraints:**

- <code>n == grid.length</code>
- <code>m == grid[i].length</code>
- <code>1 <= n, m <= 500</code>
- <code>grid[i][j]</code> is either <code>0</code>, <code>1</code> or <code>2</code>.

## Solution

Conceptually, the solution to this problem is not that hard. For every valid segment start, we have 4 directions to start in.
At each step, given that it is in the sequence, we can either continue in that direction or change direction once and follow
that sequence till its end. We need to capture the segment lengths of all those alternatives as well as the one for no 
turns. The maximum of all of those is the longest segment starting from the selected position. We can repeat this for every
segment start to get the longest overall segment. 

The time complexity of this solution is `O(m*n*max(m,n))`. The `m*n` is obvious due to the matrix. However, since the
length of the segment is bound by `max(m,n)` and we iterate this as well, we need to add it to the complexity. Addtionally,
we can make use of that bound for some optimization, i.e. if we reached the bound, we can end the calculation.