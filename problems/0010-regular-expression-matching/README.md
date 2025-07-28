# [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/description/)

Given an input string <code>s</code>and a pattern <code>p</code>, implement regular expression matching with support for <code>'.'</code> and <code>'\*'</code> where:

- <code>'.'</code> Matches any single character.​​​​
- <code>'\*'</code> Matches zero or more of the preceding element.

The matching should cover the **entire** input string (not partial).

**Example 1:**

```
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

**Example 2:**

```
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

**Example 3:**

```
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

**Constraints:**

- <code>1 <= s.length<= 20</code>
- <code>1 <= p.length<= 20</code>
- <code>s</code> contains only lowercase English letters.
- <code>p</code> contains only lowercase English letters, <code>'.'</code>, and<code>'\*'</code>.
- It is guaranteed for each appearance of the character <code>'\*'</code>, there will be a previous valid character to match.
