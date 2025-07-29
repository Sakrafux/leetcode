# [2411. Smallest Subarrays With Maximum Bitwise OR](https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/description/?envType=daily-question&envId=2025-07-29)

You are given a **0-indexed** array <code>nums</code> of length <code>n</code>, consisting of non-negative integers. For each index <code>i</code> from <code>0</code> to <code>n - 1</code>, you must determine the size of the **minimum sized** non-empty subarray of <code>nums</code> starting at <code>i</code> (**inclusive** ) that has the **maximum** possible **bitwise OR** .

- In other words, let <code>B<sub>ij</sub></code> be the bitwise OR of the subarray <code>nums[i...j]</code>. You need to find the smallest subarray starting at <code>i</code>, such that bitwise OR of this subarray is equal to <code>max(B<sub>ik</sub>)</code> where <code>i <= k <= n - 1</code>.

The bitwise OR of an array is the bitwise OR of all the numbers in it.

Return an integer array <code>answer</code> of size <code>n</code> where <code>answer[i]</code> is the length of the **minimum** sized subarray starting at <code>i</code> with **maximum** bitwise OR.

A **subarray** is a contiguous non-empty sequence of elements within an array.

**Example 1:**

```
Input: nums = [1,0,2,1,3]
Output: [3,3,2,2,1]
Explanation:
The maximum possible bitwise OR starting at any index is 3.
- Starting at index 0, the shortest subarray that yields it is [1,0,2].
- Starting at index 1, the shortest subarray that yields the maximum bitwise OR is [0,2,1].
- Starting at index 2, the shortest subarray that yields the maximum bitwise OR is [2,1].
- Starting at index 3, the shortest subarray that yields the maximum bitwise OR is [1,3].
- Starting at index 4, the shortest subarray that yields the maximum bitwise OR is [3].
Therefore, we return [3,3,2,2,1].
```

**Example 2:**

```
Input: nums = [1,2]
Output: [2,1]
**Explanation:
** Starting at index 0, the shortest subarray that yields the maximum bitwise OR is of length 2.
Starting at index 1, the shortest subarray that yields the maximum bitwise OR is of length 1.
Therefore, we return [2,1].
```

**Constraints:**

- <code>n == nums.length</code>
- <code>1 <= n <= 10^5</code>
- <code>0 <= nums[i] <= 10^9</code>

## Solution

For the problem, we want to approach each bit as its own problem. We therefore keep track of the last occurence for each bit. To automatically capture the last occurence, we iterate the numbers in reverse and then try for each bit if it occurs. Additionally, we calculate for the number in question the maximum distance to all its bits. By going in reverse we also take care of the fact that later subsets may have a smaller maximum bitwise OR. This is `O(n)`.

An alternative approach tries to keep track of distance to a specific bit for each number by using maps of int arrays where the keys are the specific bit numbers (i.e. 100 -> 4). Then by iterating all the arrays at a given index and taking the maximum value among them is the result for the index. This is in `O(n*log(m))` where `m <= 30` (as it represent the amount of bits in the maximum bitwise OR).

A simple brute force approach would simply iterate once to calculate the maxium bitwise OR and then iterate over the numbers in a nested fashion until the maximum has been reached for the outer index. This is obviously `O(n²)`
