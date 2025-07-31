# [18. 4Sum](https://leetcode.com/problems/4sum/description/)

Given an array <code>nums</code> of <code>n</code> integers, return an array of all the **unique** quadruplets <code>[nums[a], nums[b], nums[c], nums[d]]</code> such that:

- <code>0 <= a, b, c, d< n</code>
- <code>a</code>, <code>b</code>, <code>c</code>, and <code>d</code> are **distinct** .
- <code>nums[a] + nums[b] + nums[c] + nums[d] == target</code>

You may return the answer in **any order** .

**Example 1:**

```
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

**Example 2:**

```
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
```

**Constraints:**

- <code>1 <= nums.length <= 200</code>
- <code>-10^9 <= nums[i] <= 10^9</code>
- <code>-10^9 <= target <= 10^9</code>

## Solution

First, we sort the array so we know the ordering of the numbers. Then we need 2 nested loops to fix 2 numbers and reduce the problem to 2sum. Then we use 2 pointers to the left and right of the remaining array space, calculate the sum of all selected numbers and move the pointers accordingly depending on whether the sum is smaller or larger than the target. The left pointer rightwards if the sum is smaller, the right pointer leftwards if it is bigger. All the while we need to skip duplicate numbers.
