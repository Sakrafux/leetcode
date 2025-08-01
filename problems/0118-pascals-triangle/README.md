# [118. Pascal's Triangle](https://leetcode.com/problems/pascals-triangle/description/?envType=daily-question&envId=2025-08-01)

Given an integer <code>numRows</code>, return the first numRows of **Pascal's triangle** .

In **Pascal's triangle** , each number is the sum of the two numbers directly above it as shown:
<img alt="" src="https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif" style="height: 240px; width: 260px;">

**Example 1:**

```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

**Example 2:**

```
Input: numRows = 1
Output: [[1]]
```

**Constraints:**

- <code>1 <= numRows <= 30</code>

## Solution

We need to iterate `numRows` times to create new rows. While the first row is special, any other row is simply based on the previous row. It has 1 element more than the previous row. We build it with related values from the previous row, while taking care to respect the boundaries.
