# [2749. Minimum Operations to Make the Integer Zero](https://leetcode.com/problems/minimum-operations-to-make-the-integer-zero/description/?envType=daily-question&envId=2025-09-05)

You are given two integers <code>num1</code> and <code>num2</code>.

In one operation, you can choose integer <code>i</code> in the range <code>[0, 60]</code> and subtract <code>2^i + num2</code> from <code>num1</code>.

Return the integer denoting the **minimum**  number of operations needed to make <code>num1</code> equal to <code>0</code>.

If it is impossible to make <code>num1</code> equal to <code>0</code>, return <code>-1</code>.

**Example 1:**

```
Input: num1 = 3, num2 = -2
Output: 3
Explanation: We can make 3 equal to 0 with the following operations:
- We choose i = 2 and subtract 2^2 + (-2) from 3, 3 - (4 + (-2)) = 1.
- We choose i = 2 and subtract 2^2+ (-2) from 1, 1 - (4 + (-2)) = -1.
- We choose i = 0 and subtract 2^0+ (-2) from -1, (-1) - (1 + (-2)) = 0.
It can be proven, that 3 is the minimum number of operations that we need to perform.
```

**Example 2:**

```
Input: num1 = 5, num2 = 7
Output: -1
Explanation: It can be proven, that it is impossible to make 5 equal to 0 with the given operation.
```

**Constraints:**

- <code>1 <= num1 <= 10^9</code>
- <code>-10^9<= num2 <= 10^9</code>

## Solution

We need to iteratively test whether we have enough operations to solve the equation. For this we need to consider the fixed
part, i.e. `x = num1 - k*num2`, with `k` number of operations, which we then need to reach with `k` substractions of `2^i`. 
In other words, we need a high enough `k` to service all 1-bits in `x`. If at any point `x < k`, then we can't reach a solution.

The time complexity of this is `O(1)`, since the iterations are hard-bound by the input constraints.