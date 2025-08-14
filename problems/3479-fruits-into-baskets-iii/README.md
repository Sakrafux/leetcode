# [3479. Fruits Into Baskets III](https://leetcode.com/problems/fruits-into-baskets-iii/description/?envType=daily-question&envId=2025-08-06)

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
- <code>1 <= n <= 10^5</code>
- <code>1 <= fruits[i], baskets[i] <= 10^9</code>

## Solution

This is the harder version of the same problem as 3477 and thus asks for better time complexity, which invalidates
the brute force, and other, solutions of `O(n²)`. The proper approach uses segment trees, as they allow us to 
communicate multiple pieces of information at once. In this case, both the index of the basket array, which is important,
since we want to take the first basket available, while also preserving its capacity.

Segment trees can be represented as arrays, where the range information is implicit, and we can use them similar to 
a max heap. First, we create the segment tree with the index of the basket array as the range and its capacity as the value.
We can then traverse the tree left-first to find the left-most basket of sufficient capacity, which we then use up by 
setting its capacity to 0 and updating the tree accordingly. Since the parent node always contains the largest capacity,
we know instantly whether we can place a given fruit or not.

Since all the operations involved with (segment) trees are at most `O(n*log(n))`, this is our time complexity.