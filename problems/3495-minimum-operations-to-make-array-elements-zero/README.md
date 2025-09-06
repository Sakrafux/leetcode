# [3495. Minimum Operations to Make Array Elements Zero](https://leetcode.com/problems/minimum-operations-to-make-array-elements-zero/description/?envType=daily-question&envId=2025-09-06)

You are given a 2D array <code>queries</code>, where <code>queries[i]</code> is of the form <code>[l, r]</code>. Each <code>queries[i]</code> defines an array of integers <code>nums</code> consisting of elements ranging from <code>l</code> to <code>r</code>, both **inclusive** .

In one operation, you can:

- Select two integers <code>a</code> and <code>b</code> from the array.
- Replace them with <code>floor(a / 4)</code> and <code>floor(b / 4)</code>.

Your task is to determine the **minimum**  number of operations required to reduce all elements of the array to zero for each query. Return the sum of the results for all queries.

**Example 1:**

<div class="example-block">
Input: queries = [[1,2],[2,4]]

Output: 3

Explanation:

For <code>queries[0]</code>:

- The initial array is <code>nums = [1, 2]</code>.
- In the first operation, select <code>nums[0]</code> and <code>nums[1]</code>. The array becomes <code>[0, 0]</code>.
- The minimum number of operations required is 1.

For <code>queries[1]</code>:

- The initial array is <code>nums = [2, 3, 4]</code>.
- In the first operation, select <code>nums[0]</code> and <code>nums[2]</code>. The array becomes <code>[0, 3, 1]</code>.
- In the second operation, select <code>nums[1]</code> and <code>nums[2]</code>. The array becomes <code>[0, 0, 0]</code>.
- The minimum number of operations required is 2.

The output is <code>1 + 2 = 3</code>.

**Example 2:**

<div class="example-block">
Input: queries = [[2,6]]

Output: 4

Explanation:

For <code>queries[0]</code>:

- The initial array is <code>nums = [2, 3, 4, 5, 6]</code>.
- In the first operation, select <code>nums[0]</code> and <code>nums[3]</code>. The array becomes <code>[0, 3, 4, 1, 6]</code>.
- In the second operation, select <code>nums[2]</code> and <code>nums[4]</code>. The array becomes <code>[0, 3, 1, 1, 1]</code>.
- In the third operation, select <code>nums[1]</code> and <code>nums[2]</code>. The array becomes <code>[0, 0, 0, 1, 1]</code>.
- In the fourth operation, select <code>nums[3]</code> and <code>nums[4]</code>. The array becomes <code>[0, 0, 0, 0, 0]</code>.
- The minimum number of operations required is 4.

The output is 4.

**Constraints:**

- <code>1 <= queries.length <= 10^5</code>
- <code>queries[i].length == 2</code>
- <code>queries[i] == [l, r]</code>
- <code>1 <= l < r <= 10^9</code>

## Solution

This problem is quite tricky when considering efficient solutions. For each query, we need to know how many reduction 
operations are necessary, which is itself a simple divide-by-2 function of the total single element reductions. I.e. 
the amount of times we need to divide by 4. For this purpose, we can separate the range implied by a query into blocks,
whose elements require the same amount of operations. As this is based on the bits, we must iterate all possible
bit-lengths, which has a fixed upper boundary based on the input constraints. This means our query array number range
is separated into blocks of `[2^{k-1}, 2^k-1]`, which all require the same amount of operations to reduce to 0. We simply
need to repeat this for all blocks that have an overlap with our query number range and calculate the total single element
operations this way. This we can then add to the total result after correcting it for the fact that our operations affect
2 elements at once.

The time complexity of this is `O(n)` given `n` queries and the fact that checking a query is `O(31)` based on our input
constraints.

An example would be `[[3,10]]`:

Blocks:
- `k=1` : `[1,1]` -> 1 operation
- `k=2` : `[2,3]` -> 1 operation
- `k=3` : `[4,7]` -> 2 operations
- `k=4` : `[8,15]` -> 2 operations

The overlap is:
- `k=2` : `[3]` -> `1*1=1` operations
- `k=3` : `[4,7]` -> `4*2=8` operations
- `k=4` : `[8,10]` -> `3*2=6` operations

Correcting for double operations we get `(1+8+6 + 1) / 2 = 8`