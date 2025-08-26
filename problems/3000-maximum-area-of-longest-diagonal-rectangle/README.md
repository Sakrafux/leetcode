# [3000. Maximum Area of Longest Diagonal Rectangle](https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/description/?envType=daily-question&envId=2025-08-26)

You are given a 2D **0-indexed ** integer array <code>dimensions</code>.

For all indices <code>i</code>, <code>0 <= i < dimensions.length</code>, <code>dimensions[i][0]</code> represents the length and <code>dimensions[i][1]</code> represents the width of the rectangle <code>i</code>.

Return the **area**  of the rectangle having the **longest**  diagonal. If there are multiple rectangles with the longest diagonal, return the area of the rectangle having the **maximum**  area.

**Example 1:**

```
Input: dimensions = [[9,3],[8,6]]
Output: 48
Explanation: 
For index = 0, length = 9 and width = 3. Diagonal length = sqrt(9 * 9 + 3 * 3) = sqrt(90) ≈ 9.487.
For index = 1, length = 8 and width = 6. Diagonal length = sqrt(8 * 8 + 6 * 6) = sqrt(100) = 10.
So, the rectangle at index 1 has a greater diagonal length therefore we return area = 8 * 6 = 48.
```

**Example 2:**

```
Input: dimensions = [[3,4],[4,3]]
Output: 12
Explanation: Length of diagonal is the same for both which is 5, so maximum area = 12.
```

**Constraints:**

- <code>1 <= dimensions.length <= 100</code>
- <code>dimensions[i].length == 2</code>
- <code>1 <= dimensions[i][0], dimensions[i][1] <= 100</code>

## Solution

The very simple approach simply iterates all given dimensions and calculates their diagonal. It should be noted that
it is not necessary to take the square root, since it is simply a monotone transformation, and we are not interested
in the diagonal itself, but rather its ordering. Then we calculate the area, if the diagonal is longer or equal, and in
the case of the latter, take the greater area.

The time complexity is `O(n)`, since we iterate all entries of the `dimensions` array once.