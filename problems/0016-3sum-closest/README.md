# [16. 3Sum Closest](https://leetcode.com/problems/3sum-closest/description/)

Given an integer array <code>nums</code> of length <code>n</code> and an integer <code>target</code>, find three integers in <code>nums</code> such that the sum is closest to <code>target</code>.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

**Example 1:**

```
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
```

**Example 2:**

```
Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
```

**Constraints:**

- <code>3 <= nums.length <= 500</code>
- <code>-1000 <= nums[i] <= 1000</code>
- <code>-10^4 <= target <= 10^4</code>

## Solution

We sort the array so we know the ordering of the numbers and then fix one number through an loop. Inside, we use 2 pointers, one to each end of the remaining array space, to select number combinations. If the chosen combination is closer than the current best, we have found a better solution. Due to the sorted array, we know that moving the left pointer rightwards increases the solution. Moving the right pointer leftwards decreases it. Depending on whether the solution is larger or smaller than the target, we move the pointers accordingly until we either run out of numbers or reach the target solution.
