# [8. String to Integer (atoi)](https://leetcode.com/problems/string-to-integer-atoi/description/)

Implement the <code>myAtoi(string s)</code> function, which converts a string to a 32-bit signed integer.

The algorithm for <code>myAtoi(string s)</code> is as follows:

- **Whitespace** : Ignore any leading whitespace (<code>" "</code>).
- **Signedness** : Determine the sign by checking if the next character is <code>'-'</code> or <code>'+'</code>, assuming positivity if neither present.
- **Conversion** : Read the integer by skipping leading zerosuntil a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
- **Rounding** : If the integer is out of the 32-bit signed integer range <code>[-2^31, 2^31 - 1]</code>, then round the integer to remain in the range. Specifically, integers less than <code>-2^31</code> should be rounded to <code>-2^31</code>, and integers greater than <code>2^31 - 1</code> should be rounded to <code>2^31 - 1</code>.

Return the integer as the final result.

**Example 1:**

<div class="example-block">
Input: s = "42"

Output: 42

Explanation:

```
The underlined characters are what is read in and the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
```

**Example 2:**

<div class="example-block">
Input: s = " -042"

Output: -42

Explanation:

```
Step 1: "   -042" (leading whitespace is read and ignored)
            ^
Step 2: "   -042" ('-' is read, so the result should be negative)
             ^
Step 3: "   -042" ("042" is read in, leading zeros ignored in the result)
               ^
```

**Example 3:**

<div class="example-block">
Input: s = "1337c0d3"

Output: 1337

Explanation:

```
Step 1: "1337c0d3" (no characters read because there is no leading whitespace)
         ^
Step 2: "1337c0d3" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "1337c0d3" ("1337" is read in; reading stops because the next character is a non-digit)
             ^
```

**Example 4:**

<div class="example-block">
Input: s = "0-1"

Output: 0

Explanation:

```
Step 1: "0-1" (no characters read because there is no leading whitespace)
         ^
Step 2: "0-1" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "0-1" ("0" is read in; reading stops because the next character is a non-digit)
          ^
```

**Example 5:**

<div class="example-block">
Input: s = "words and 987"

Output: 0

Explanation:

Reading stops at the first non-digit character 'w'.

**Constraints:**

- <code>0 <= s.length <= 200</code>
- <code>s</code> consists of English letters (lower-case and upper-case), digits (<code>0-9</code>), <code>' '</code>, <code>'+'</code>, <code>'-'</code>, and <code>'.'</code>.

## Solution

The solution is rather straightforward, with the only issue really being the possible exceptions, i.e. an empty string or a blank string (only whitespaces). Otherwise, simply skip the whitespaces, read the sign if possible. Then read the remaining string, ending on non-numeric characters, skip '0' as long as the result is 0, i.e. leading zeroes, and add the digit to the result, keeping in mind that '0' is actually an ASCII code and not 0. Since we are not constrained to int32 on the system, we can simply add the digit and then check for int32 boundaries in order to round.
