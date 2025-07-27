# [2210. Count Hills and Valleys in an Array](https://leetcode.com/problems/count-hills-and-valleys-in-an-array/description/?envType=daily-question&envId=2025-07-27)

You are given a **0-indexed** integer array <code>nums</code>. An index <code>i</code> is part of a **hill** in <code>nums</code> if the closest non-equal neighbors of <code>i</code> are smaller than <code>nums[i]</code>. Similarly, an index <code>i</code> is part of a **valley** in <code>nums</code> if the closest non-equal neighbors of <code>i</code> are larger than <code>nums[i]</code>. Adjacent indices <code>i</code> and <code>j</code> are part of the **same** hill or valley if <code>nums[i] == nums[j]</code>.

Note that for an index to be part of a hill or valley, it must have a non-equal neighbor on **both** the left and right of the index.

Return <i>the number of hills and valleys in </i><code>nums</code>.

**Example 1:**

```
Input: nums = [2,4,1,1,6,5]
Output: 3
Explanation:
At index 0: There is no non-equal neighbor of 2 on the left, so index 0 is neither a hill nor a valley.
At index 1: The closest non-equal neighbors of 4 are 2 and 1. Since 4 > 2 and 4 > 1, index 1 is a hill.
At index 2: The closest non-equal neighbors of 1 are 4 and 6. Since 1 < 4 and 1 < 6, index 2 is a valley.
At index 3: The closest non-equal neighbors of 1 are 4 and 6. Since 1 < 4 and 1 < 6, index 3 is a valley, but note that it is part of the same valley as index 2.
At index 4: The closest non-equal neighbors of 6 are 1 and 5. Since 6 > 1 and 6 > 5, index 4 is a hill.
At index 5: There is no non-equal neighbor of 5 on the right, so index 5 is neither a hill nor a valley.
There are 3 hills and valleys so we return 3.
```

**Example 2:**

```
Input: nums = [6,6,5,5,4,1]
Output: 0
Explanation:
At index 0: There is no non-equal neighbor of 6 on the left, so index 0 is neither a hill nor a valley.
At index 1: There is no non-equal neighbor of 6 on the left, so index 1 is neither a hill nor a valley.
At index 2: The closest non-equal neighbors of 5 are 6 and 4. Since 5 < 6 and 5 > 4, index 2 is neither a hill nor a valley.
At index 3: The closest non-equal neighbors of 5 are 6 and 4. Since 5 < 6 and 5 > 4, index 3 is neither a hill nor a valley.
At index 4: The closest non-equal neighbors of 4 are 5 and 1. Since 4 < 5 and 4 > 1, index 4 is neither a hill nor a valley.
At index 5: There is no non-equal neighbor of 1 on the right, so index 5 is neither a hill nor a valley.
There are 0 hills and valleys so we return 0.
```

**Constraints:**

- <code>3 <= nums.length <= 100</code>
- <code>1 <= nums[i] <= 100</code>

## Solution

We only need to iterate once over the array, thus O(n). We have effectivly 3 pointers, one to the left, the middle, and the right value, with the middle value being our iterator and excluding the first and last element.

We iterate to move the middle and right forwards until the middle is different to both left and right. Only in this situation do we check for hill or valley and reassign left to the current value of middle.
