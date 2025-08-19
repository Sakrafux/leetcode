# [2348. Number of Zero-Filled Subarrays](https://leetcode.com/problems/number-of-zero-filled-subarrays/description/?envType=daily-question&envId=2025-08-19)

Given an integer array <code>nums</code>, return the number of **subarrays**  filled with <code>0</code>.

A **subarray**  is a contiguous non-empty sequence of elements within an array.

**Example 1:**

```
Input: nums = [1,3,0,0,2,0,0,4]
Output: 6
Explanation: 
There are 4 occurrences of [0] as a subarray.
There are 2 occurrences of [0,0] as a subarray.
There is no occurrence of a subarray with a size more than 2 filled with 0. Therefore, we return 6.
```

**Example 2:** 

```
Input: nums = [0,0,0,2,0,0]
Output: 9
**Explanation:
** There are 5 occurrences of [0] as a subarray.
There are 3 occurrences of [0,0] as a subarray.
There is 1 occurrence of [0,0,0] as a subarray.
There is no occurrence of a subarray with a size more than 3 filled with 0. Therefore, we return 9.
```

**Example 3:** 

```
Input: nums = [2,10,2019]
Output: 0
Explanation: There is no subarray filled with 0. Therefore, we return 0.
```

**Constraints:** 

- <code>1 <= nums.length <= 10^5</code>
- <code>-10^9 <= nums[i] <= 10^9</code>

## Solution

We need to iterate through the array and continuously keep track of both our total sum and of the consecutive zeros we
have seen so far. Because each new number we look at, if it is 0, can create a new subarray with each additional previous
zero. E.g. if we have the 3rd zero in a row, it enables both [0, 0] and [0, 0, 0]. In addition, every zero allows for 
a single element subarray.

The time complexity of this is `O(n)` as we are iterating the array a single time.