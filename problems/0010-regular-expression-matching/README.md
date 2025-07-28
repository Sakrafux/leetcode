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

## Solution

This is essentially a dynamic programming problem where we want to progressively parse the target string and build up the possible parsing combination until the final element (i.e. both strings completely parsed) tells us, whether the match is successful.

The dynamic programming table is `length s + 1` rows and `length p + 1` columns, with the 0 row/column representing empty strings. The first table element 0/0 is thus true, because empty string matches empty pattern. We then populate the first row with repeating occurences that can match empty strings, i.e. any leading `*` patterns.

Then we work row by row, i.e. character by character of the target string, and always iterate the entire pattern. If the characters match, we simply inherit the result of the matching sofar. If the iterated character of the pattern is `*`, it could either be zero occurences or multiple occurences.

An example:

```
s = "abaaab"
p = "b*ab.*ab"

      0 1 2 3 4 5 6 7 8
        b * a b . * a b
    ___________________
0   | 1 0 1 0 0 0 0 0 0
1 a | 0 0 0 1 0 0 0 0 0
2 b | 0 0 0 0 1 0 1 0 0
3 a | 0 0 0 0 0 1 1 1 0
4 a | 0 0 0 0 0 0 1 1 0
5 a | 0 0 0 0 0 0 1 1 0
6 b | 0 0 0 0 0 0 1 0 1
```

In 0/2 and 2/6 we can see how `*` accounts for no occurences. Whereas in 3/6, 4/6, 5/6, and 6/6 we can see multiple occurences being accounted for.
