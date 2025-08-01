# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/description/)

Given a string <code>s</code> containing just the characters <code>'('</code>, <code>')'</code>, <code>'{'</code>, <code>'}'</code>, <code>'['</code> and <code>']'</code>, determine if the input string is valid.

An input string is valid if:

- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

**Example 1:**

<div class="example-block">
Input: s = "()"

Output: true

**Example 2:**

<div class="example-block">
Input: s = "()[]{}"

Output: true

**Example 3:**

<div class="example-block">
Input: s = "(]"

Output: false

**Example 4:**

<div class="example-block">
Input: s = "([])"

Output: true

**Example 5:**

<div class="example-block">
Input: s = "([)]"

Output: false

**Constraints:**

- <code>1 <= s.length <= 10^4</code>
- <code>s</code> consists of parentheses only <code>'()[]{}'</code>.

## Solution

The solution is to simply use a stack. All opening brackets are put onto the stack and closing brackets must fit the uppermost element of the stack. If anything doesn't fit or the stack doesn't empty, it can't be valid.
