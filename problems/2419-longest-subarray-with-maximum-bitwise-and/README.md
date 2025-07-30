# [2419. Longest Subarray With Maximum Bitwise AND](https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/description/?envType=daily-question&envId=2025-07-30)

You are given an integer array <code>nums</code> of size <code>n</code>.

Consider a **non-empty** subarray from <code>nums</code> that has the **maximum** possible **bitwise AND** .

- In other words, let <code>k</code> be the maximum value of the bitwise AND of **any** subarray of <code>nums</code>. Then, only subarrays with a bitwise AND equal to <code>k</code> should be considered.

Return the length of the **longest** such subarray.

The bitwise AND of an array is the bitwise AND of all the numbers in it.

A **subarray** is a contiguous sequence of elements within an array.

**Example 1:**

```
Input: nums = [1,2,3,3,2,2]
Output: 2
Explanation:
The maximum possible bitwise AND of a subarray is 3.
The longest subarray with that value is [3,3], so we return 2.
```

**Example 2:**

```
Input: nums = [1,2,3,4]
Output: 1
Explanation:
The maximum possible bitwise AND of a subarray is 4.
The longest subarray with that value is [4], so we return 1.
```

**Constraints:**

- <code>1 <= nums.length <= 10^5</code>
- <code>1 <= nums[i] <= 10^6</code>

## Solution

We can make use of the fact that the largest bitwise AND will always be the largest number only. Therefore, we only need to find the largest subarray only containing the largest number. We have to take care of the case that there may be multiple, separated subarrays containing the largest number that must not be added together
