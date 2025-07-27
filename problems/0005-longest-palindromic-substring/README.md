# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/description/)

Given a string <code>s</code>, return the longest <button type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:rs:" data-state="closed" class="">palindromic</button> <button type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:rt:" data-state="closed" class="">substring</button> in <code>s</code>.

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"
```

**Constraints:**

- <code>1 <= s.length <= 1000</code>
- <code>s</code> consist of only digits and English letters.

## Solution

While there are multiple approaches to tackling this problem, I opted for solving it via expansion from the centers.

For this, we iterate over the string with the index indicating the current center. Then, in a loop, expand symetrically until we hit the boundary. Each step we check whether the character on the left and right side are the same and if they are, we have (still) a palindrome and can check whether it is larger than what we already know. If it is not a palindrome, we can instantly break, because it will never become one again.

This covers all the cases with a single center, but there may be even-length palindromes, i.e. the center is between two characters. For this we simply need to adjust the logic a little to not use a single center but two adjacent characters and then apply this loop also.

As we are looping twice over the array, we have a complexity of `O(n^2)`.
