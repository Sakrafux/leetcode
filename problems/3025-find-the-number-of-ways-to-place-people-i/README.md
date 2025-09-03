# [3025. Find the Number of Ways to Place People I](https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/description/?envType=daily-question&envId=2025-09-02)

You are given a 2D array <code>points</code> of size <code>n x 2</code> representing integer coordinates of some points on a 2D plane, where <code>points[i] = [x<sub>i</sub>, y<sub>i</sub>]</code>.

Count the number of pairs of points <code>(A, B)</code>, where

- <code>A</code> is on the **upper left**  side of <code>B</code>, and
- there are no other points in the rectangle (or line) they make (**including the border** ).

Return the count.

**Example 1:**

<div class="example-block">
Input: points = [[1,1],[2,2],[3,3]]

Output: 0

Explanation:

<img src="https://assets.leetcode.com/uploads/2024/01/04/example1alicebob.png" style="width: 427px; height: 350px;">

There is no way to choose <code>A</code> and <code>B</code> so <code>A</code> is on the upper left side of <code>B</code>.

**Example 2:**

<div class="example-block">
Input: points = [[6,2],[4,4],[2,6]]

Output: 2

Explanation:

<img height="365" width="1321" src="https://assets.leetcode.com/uploads/2024/06/25/t2.jpg">

- The left one is the pair <code>(points[1], points[0])</code>, where <code>points[1]</code> is on the upper left side of <code>points[0]</code> and the rectangle is empty.
- The middle one is the pair <code>(points[2], points[1])</code>, same as the left one it is a valid pair.
- The right one is the pair <code>(points[2], points[0])</code>, where <code>points[2]</code> is on the upper left side of <code>points[0]</code>, but <code>points[1]</code> is inside the rectangle so it's not a valid pair.

**Example 3:**

<div class="example-block">
Input: points = [[3,1],[1,3],[1,1]]

Output: 2

Explanation:

<img src="https://assets.leetcode.com/uploads/2024/06/25/t3.jpg" style="width: 1269px; height: 350px;">

- The left one is the pair <code>(points[2], points[0])</code>, where <code>points[2]</code> is on the upper left side of <code>points[0]</code> and there are no other points on the line they form. Note that it is a valid state when the two points form a line.
- The middle one is the pair <code>(points[1], points[2])</code>, it is a valid pair same as the left one.
- The right one is the pair <code>(points[1], points[0])</code>, it is not a valid pair as <code>points[2]</code> is on the border of the rectangle.

**Constraints:**

- <code>2 <= n <= 50</code>
- <code>points[i].length == 2</code>
- <code>0 <= points[i][0], points[i][1] <= 50</code>
- All <code>points[i]</code> are distinct.

## Solution

To solve this problem, we first need to sort the pairs we are given from top-left to bottom-right. More specifically, 
we sort by the x-axis (ascending) and then y-axis (descending). This ensures that only an earlier point can be the top-left
of a later point, allowing us to iterate with confidence, checking all pairs of points. We also need to take care to not
trap any points within a rectangle, which we can do by keeping track of boundaries, i.e. the x- and y-coordinates of the
last valid bottom-right. This works because of our sorting.

The time complexity of this is `O(n²)`, because the sorting in `O(n*log(n))` is overshadowed by the pair-wise comparison
in `O(n²)` due to the nested iteration.

A much simpler solution with worse time complexity is simply brute-forcing all combinations. As this would entail iterating
once for point A, once nested for point B, and once more nested to check whether any other point lies in between, this
is `O(n³)`.