# [2044. Count Number of Maximum Bitwise-OR Subsets](https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/description/?envType=daily-question&envId=2025-07-28)

Given an integer array <code>nums</code>, find the **maximum** possible **bitwise OR** of a subset of <code>nums</code> and return the **number of different non-empty subsets** with the maximum bitwise OR.

An array <code>a</code> is a **subset** of an array <code>b</code> if <code>a</code> can be obtained from <code>b</code> by deleting some (possibly zero) elements of <code>b</code>. Two subsets are considered **different** if the indices of the elements chosen are different.

The bitwise OR of an array <code>a</code> is equal to <code>a[0] **OR** a[1] **OR** ... **OR** a[a.length - 1]</code> (**0-indexed** ).

**Example 1:**

```
Input: nums = [3,1]
Output: 2
Explanation: The maximum possible bitwise OR of a subset is 3. There are 2 subsets with a bitwise OR of 3:
- [3]
- [3,1]
```

**Example 2:**

```
Input: nums = [2,2,2]
Output: 7
Explanation: All non-empty subsets of [2,2,2] have a bitwise OR of 2. There are 2^3 - 1 = 7 total subsets.
```

**Example 3:**

````
Input: nums = [3,2,1,5]
Output: 6
Explanation: The maximum possible bitwise OR of a subset is 7. There are 6 subsets with a bitwise OR of 7:
- [3,5]
- [3,1,5]
- [3,2,5]
- [3,2,1,5]
- [2,5]
- [2,1,5]```

**Constraints:**

- <code>1 <= nums.length <= 16</code>
- <code>1 <= nums[i] <= 10^5</code>
````

## Solution

The simple brute force solution just tries to generate each subset combination and forms the relevant bitwise OR and checks against the current maximum value. This is `O(2^n)`.

A more sophisticated approach makes use of the properties of the bitwise OR. For this, we initialize an array for our possible OR-values, i.e. 2^17-1. Then we iterate over the available numbers and then iterate based on the current maximum value backwards (backwards so as not to falsify the ongoing iteration, a cloned array should also work). This counts all combinations a sum can be reached through bitwise-OR.
