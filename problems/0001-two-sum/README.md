# [1. Two Sum](https://leetcode.com/problems/two-sum/description/)

Given an array of integers <code>nums</code>and an integer <code>target</code>, return indices of the two numbers such that they add up to <code>target</code>.

You may assume that each input would have **exactly one solution** , and you may not use the same element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- <code>2 <= nums.length <= 10^4</code>
- <code>-10^9 <= nums[i] <= 10^9</code>
- <code>-10^9 <= target <= 10^9</code>
- **Only one valid answer exists.**

**Follow-up:** Can you come up with an algorithm that is less than <code>O(n^2)</code>time complexity?

## Solution

### Brute-Force - O(n²)

The brute-force solution tries each item combination by iterating twice over the array in a nested fashion. Since the constraints ensure a solution, we only need to run until the first success.

### Hash-Map - O(n)

The optimized solution makes use of the fact that we know 2 of 3 values. I.e. we need `x + y = target`, but as `target` is fixed and by iterating over the array, in each iteration `x` is effectively fixed, we only need `y`. We can now keep track of the values we already iterated over in a hash-map and look-up if `y = target - x` is already present in the map, which would then be the solution.
