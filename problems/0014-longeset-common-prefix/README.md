# [14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/description/)

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string <code>""</code>.

**Example 1:**

```
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

**Example 2:**

```
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

**Constraints:**

- <code>1 <= strs.length <= 200</code>
- <code>0 <= strs[i].length <= 200</code>
- <code>strs[i]</code> consists of only lowercase English letters if it is non-empty.

## Solution

Iteratively increase a position pointer, take the current character from the first string (if possible) and then iterate through all strings to check if they have the same character at that position. If not, return the current result, otherwise add the character to the result.

Time complexity is `O(n*m)` where `n` is the number of strings and `m` the maximum length of a single string.
