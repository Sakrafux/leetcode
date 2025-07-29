# [15. 3Sum](https://leetcode.com/problems/3sum/description/)

Given an integer array nums, return all the triplets <code>[nums[i], nums[j], nums[k]]</code> such that <code>i != j</code>, <code>i != k</code>, and <code>j != k</code>, and <code>nums[i] + nums[j] + nums[k] == 0</code>.

Notice that the solution set must not contain duplicate triplets.

**Example 1:**

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
```

**Example 2:**

```
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
```

**Example 3:**

```
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
```

**Constraints:**

- <code>3 <= nums.length <= 3000</code>
- <code>-10^5 <= nums[i] <= 10^5</code>

## Solution

In advance, we need to sort the numbers so we can make use of assured order. Furthermore, handling 3 variables at the same time is too difficult, so we iterate over the numbers and treat the index/number of the current iteration as fixed. Now we only have 2 moving parts in 2 pointers, one left and one right of the remaining array space.

We make use of the fact that the array is ordered and therefore know that the left pointer's value must always be smaller or equal to the right pointer's value. Thus, if the sum of the 3 numbers is smaller than 0, we need to increase it, which we can achieve by moving the left pointer towards bigger numbers (i.e. rightwards). Similarly, a sum bigger than 0 can be decreased by moving the righ pointer leftwards.

Should the sum be 0, then we have found a solution and then move the left pointer rightwards until we find a different number (in order to prevent duplicate combinations) for the next iteration. If at any point, the left pointer overtakes the right, then we stop and go into the next iteration of the outer loop.

This has a time complexity of `O(n²)` as we iterated over the array once nested.

### Alternate approach

An alternate approach also utilized the fixing of one index/number using an outer loop, but then tries to find the combination using a two-sum algorithm and checking for duplicate combinations using a nested map.

Theoretically, this should also be `O(n²)`, as the outer loop is `O(n)` and the two-sum algorithm is also `O(n)` (and nested inside the loop), while maps and sorting of fixed size arrays (of size 3) are all `O(1)`. However, practically all the fixed complexity operations amount to significant time.
