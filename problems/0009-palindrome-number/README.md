# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)

The string <code>"PAYPALISHIRING"</code> is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: <code>"PAHNAPLSIIGYIR"</code>

Write the code that will take a string and make this conversion given a number of rows:

```
string convert(string s, int numRows);
```

**Example 1:**

```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

**Example 2:**

```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
```

**Example 3:**

```
Input: s = "A", numRows = 1
Output: "A"
```

**Constraints:**

- <code>1 <= s.length <= 1000</code>
- <code>s</code> consists of English letters (lower-case and upper-case), <code>','</code> and <code>'.'</code>.
- <code>1 <= numRows <= 1000</code>
