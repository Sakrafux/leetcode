# [1304. Find N Unique Integers Sum up to Zero](https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/?envType=daily-question&envId=2025-09-07)

Given an integer <code>n</code>, return **any**  array containing <code>n</code> **unique**  integers such that they add up to <code>0</code>.

**Example 1:**

```
Input: n = 5
Output: [-7,-1,1,3,4]
Explanation: These arrays also are accepted [-5,-1,1,2,3] , [-3,-1,2,-2,4].
```

**Example 2:**

```
Input: n = 3
Output: [-1,0,1]
```

**Example 3:**

```
Input: n = 1
Output: [0]
```

**Constraints:**

- <code>1 <= n <= 1000</code>

## Solution

This problem is rather simple as we have access to both positive and negative integers. So we can simply add `n/2` pairs of
positive and negative numbers with the same absolute value each. As a provision for odd `n`, we can add 0.

The time complexity of this is `O(n)`.