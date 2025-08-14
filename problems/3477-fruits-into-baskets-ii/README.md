# [3477. Fruits Into Baskets II](https://leetcode.com/problems/fruits-into-baskets-ii/description/?envType=daily-question&envId=2025-08-05)

You are given two arrays of integers, <code>fruits</code> and <code>baskets</code>, each of length <code>n</code>, where <code>fruits[i]</code> represents the **quantity** of the <code>i^th</code> type of fruit, and <code>baskets[j]</code> represents the **capacity** of the <code>j^th</code> basket.

From left to right, place the fruits according to these rules:

- Each fruit type must be placed in the **leftmost available basket** with a capacity **greater than or equal** to the quantity of that fruit type.
- Each basket can hold <b>only one</b> type of fruit.
- If a fruit type <b>cannot be placed</b> in any basket, it remains <b>unplaced</b>.

Return the number of fruit types that remain unplaced after all possible allocations are made.

**Example 1:**

<div class="example-block">
Input: fruits = [4,2,5], baskets = [3,5,4]

Output: 1

Explanation:

- <code>fruits[0] = 4</code> is placed in <code>baskets[1] = 5</code>.
- <code>fruits[1] = 2</code> is placed in <code>baskets[0] = 3</code>.
- <code>fruits[2] = 5</code> cannot be placed in <code>baskets[2] = 4</code>.

Since one fruit type remains unplaced, we return 1.

**Example 2:**

<div class="example-block">
Input: fruits = [3,6,1], baskets = [6,4,7]

Output: 0

Explanation:

- <code>fruits[0] = 3</code> is placed in <code>baskets[0] = 6</code>.
- <code>fruits[1] = 6</code> cannot be placed in <code>baskets[1] = 4</code> (insufficient capacity) but can be placed in the next available basket, <code>baskets[2] = 7</code>.
- <code>fruits[2] = 1</code> is placed in <code>baskets[1] = 4</code>.

Since all fruits are successfully placed, we return 0.

**Constraints:**

- <code>n == fruits.length == baskets.length</code>
- <code>1 <= n <= 100</code>
- <code>1 <= fruits[i], baskets[i] <= 1000</code>

## Solution

Since this is the easy version of the problem, this asks for a simple approach. The simple solution is to iterate
both arrays in a nested manner, with the outer loop iterating the fruits and the inner loop the baskets. Once we 
find a fit, we simply set the capacity negative and move on to the next fruit. If we don't find a fit, it can't
be placed.

The time complexity of this is `O(n²)`.