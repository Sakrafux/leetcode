# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/description/)

Given an integer <code>x</code>, return <code>true</code> if <code>x</code> is a <button type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:rs:" data-state="closed" class="">**palindrome** </button>, and <code>false</code> otherwise.

**Example 1:**

```
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
```

**Example 2:**

```
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3:**

```
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

**Constraints:**

- <code>-2^31<= x <= 2^31- 1</code>

**Follow up:** Could you solve it without converting the integer to a string?

## Solution

Simply reverse the number and compare with the original. Additionally, catch negative numbers as false.
