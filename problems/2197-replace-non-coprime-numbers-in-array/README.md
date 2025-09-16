# [2197. Replace Non-Coprime Numbers in Array](https://leetcode.com/problems/replace-non-coprime-numbers-in-array/description/?envType=daily-question&envId=2025-09-16)

You are given an array of integers <code>nums</code>. Perform the following steps:

- Find **any**  two **adjacent**  numbers in <code>nums</code> that are **non-coprime** .
- If no such numbers are found, **stop**  the process.
- Otherwise, delete the two numbers and **replace**  them with their **LCM (Least Common Multiple)** .
- **Repeat**  this process as long as you keep finding two adjacent non-coprime numbers.

Return the **final**  modified array. It can be shown that replacing adjacent non-coprime numbers in **any**  arbitrary order will lead to the same result.

The test cases are generated such that the values in the final array are **less than or equal**  to <code>10^8</code>.

Two values <code>x</code> and <code>y</code> are **non-coprime**  if <code>GCD(x, y) > 1</code> where <code>GCD(x, y)</code> is the **Greatest Common Divisor**  of <code>x</code> and <code>y</code>.

**Example 1:**

```
Input: nums = [6,4,3,2,7,6,2]
Output: [12,7,6]
Explanation: 
- (6, 4) are non-coprime with LCM(6, 4) = 12. Now, nums = [**12** ,3,2,7,6,2].
- (12, 3) are non-coprime with LCM(12, 3) = 12. Now, nums = [**12** ,2,7,6,2].
- (12, 2) are non-coprime with LCM(12, 2) = 12. Now, nums = [**12** ,7,6,2].
- (6, 2) are non-coprime with LCM(6, 2) = 6. Now, nums = [12,7,**6** ].
There are no more adjacent non-coprime numbers in nums.
Thus, the final modified array is [12,7,6].
Note that there are other ways to obtain the same resultant array.
```

**Example 2:**

```
Input: nums = [2,2,1,1,3,3,3]
Output: [2,1,1,3]
Explanation: 
- (3, 3) are non-coprime with LCM(3, 3) = 3. Now, nums = [2,2,1,1,**3** ,3].
- (3, 3) are non-coprime with LCM(3, 3) = 3. Now, nums = [2,2,1,1,**3** ].
- (2, 2) are non-coprime with LCM(2, 2) = 2. Now, nums = [**2** ,1,1,3].
There are no more adjacent non-coprime numbers in nums.
Thus, the final modified array is [2,1,1,3].
Note that there are other ways to obtain the same resultant array.
```

**Constraints:**

- <code>1 <= nums.length <= 10^5</code>
- <code>1 <= nums[i] <= 10^5</code>
- The test cases are generated such that the values in the final array are **less than or equal**  to <code>10^8</code>.

## Solution

The optimal solution makes use of the fact that the LCM can be calculated using the GCD, and that we don't need to work
with the entire array at once. We iterate through the original numbers, append them to our result array. As long as we have
more than 1 element, we get the GCD for the last 2 elements, and if GCD > 1, remove the elements and instead append the LCM, 
thus shrinking the array by 1, otherwise we simply break the loop, because we can't shrink the array.

The time complexity of this is `O(n*log(m))`, where `n = nums.length` and `m = nums[i].size`, because the nested loop 
only executes fixed number of operations per element.