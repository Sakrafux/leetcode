# [7. Reverse Integer](https://leetcode.com/problems/reverse-integer/description/)

Given a signed 32-bit integer <code>x</code>, return <code>x</code> with its digits reversed. If reversing <code>x</code> causes the value to go outside the signed 32-bit integer range <code>[-2^31, 2^31 - 1]</code>, then return <code>0</code>.

**Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**

**Example 1:**

```
Input: x = 123
Output: 321
```

**Example 2:**

```
Input: x = -123
Output: -321
```

**Example 3:**

```
Input: x = 120
Output: 21
```

**Constraints:**

- <code>-2^31 <= x <= 2^31 - 1</code>

## Solution

The standard approach for reversing a number is rather simple, by repeatedly taking its modulo 10 for the digit, dividing the number by 10, and building the result by multiplying by 10 instead and adding the digit. However, the difficulty lies in recognizing the limits. For this purpose, we precalculate the constants that may not be crossed based on the limits, i.e. if the result is already greater/lesser than those constants then we can't proceed anymore.
