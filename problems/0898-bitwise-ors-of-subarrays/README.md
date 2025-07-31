# [898. Bitwise ORs of Subarrays](https://leetcode.com/problems/bitwise-ors-of-subarrays/description/?envType=daily-question&envId=2025-07-31)

Given an integer array <code>arr</code>, return the number of distinct bitwise ORs of all the non-empty subarrays of <code>arr</code>.

The bitwise OR of a subarray is the bitwise OR of each integer in the subarray. The bitwise OR of a subarray of one integer is that integer.

A **subarray** is a contiguous non-empty sequence of elements within an array.

**Example 1:**

```
Input: arr = [0]
Output: 1
Explanation: There is only one possible result: 0.
```

**Example 2:**

```
Input: arr = [1,1,2]
Output: 3
Explanation: The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
These yield the results 1, 1, 2, 1, 3, 3.
There are 3 unique values, so the answer is 3.
```

**Example 3:**

```
Input: arr = [1,2,4]
Output: 6
Explanation: The possible results are 1, 2, 3, 4, 6, and 7.
```

**Constraints:**

- <code>1 <= arr.length <= 5 \* 10^4</code>
- <code>0 <= arr[i] <= 10^9</code>

## Solution

We make use of the fact that bitwise OR is monotonic, i.e. `x|y >= x && x|y >= y`. Given that we are limited by 32-bit int, there is only 32 bits that can be activated.

We apply this knowledge by iterating over the numbers, adding the number to both a global result set and a current set. Adding the number represents the start of a new subarray. Then we iterate over all the values of the previous iterations set, which is heavily limited in size by the aforementioned fact, i.e. `<= 32` iterations. We construct the bitwise OR and add it to the global and the current set. After iterating the previous set, we update it to the current set.

The time complexity of this is `O(n)`, since the inner loop is `O(32) ~ O(1)`.

A brute force attempt would simply iterate twice over the array, building all combinations and save them in a set. This is `O(n²)`.
