# [12. Integer to Roman](https://leetcode.com/problems/integer-to-roman/description/)

Seven different symbols represent Roman numerals with the following values:

<table><thead><tr><th>Symbol</th><th>Value</th></tr></thead><tbody><tr><td>I</td><td>1</td></tr><tr><td>V</td><td>5</td></tr><tr><td>X</td><td>10</td></tr><tr><td>L</td><td>50</td></tr><tr><td>C</td><td>100</td></tr><tr><td>D</td><td>500</td></tr><tr><td>M</td><td>1000</td></tr></tbody></table>

Roman numerals are formed by appendingthe conversions ofdecimal place valuesfrom highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

- If the value does not start with 4 or9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
- If the value starts with 4 or 9 use the**subtractive form** representingone symbol subtracted from the following symbol, for example,4 is 1 (<code>I</code>) less than 5 (<code>V</code>): <code>IV</code>and 9 is 1 (<code>I</code>) less than 10 (<code>X</code>): <code>IX</code>.Only the following subtractive forms are used: 4 (<code>IV</code>), 9 (<code>IX</code>),40 (<code>XL</code>), 90 (<code>XC</code>), 400 (<code>CD</code>) and 900 (<code>CM</code>).
- Only powers of 10 (<code>I</code>, <code>X</code>, <code>C</code>, <code>M</code>) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5(<code>V</code>), 50 (<code>L</code>), or 500 (<code>D</code>) multiple times. If you need to append a symbol4 timesuse the **subtractive form** .

Given an integer, convert it to a Roman numeral.

**Example 1:**

<div class="example-block">
Input: num = 3749

Output: "MMMDCCXLIX"

Explanation:

```
3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
 700 = DCC as 500 (D) + 100 (C) + 100 (C)
  40 = XL as 10 (X) less of 50 (L)
   9 = IX as 1 (I) less of 10 (X)
Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
```

**Example 2:**

<div class="example-block">
Input: num = 58

Output: "LVIII"

Explanation:

```
50 = L
 8 = VIII
```

**Example 3:**

<div class="example-block">
Input: num = 1994

Output: "MCMXCIV"

Explanation:

```
1000 = M
 900 = CM
  90 = XC
   4 = IV
```

**Constraints:**

- <code>1 <= num <= 3999</code>

## Solution

Due to the definition of roman numbers, we have a limited mapping of symbols, i.e. which numbers can be represented by a single token (consisting of 1 to 2 characters). Creating all those mappings, we can then iterate over them and if the value is larger than the number represented by the mapping, we substract the mapping's value and add its token to the result string. Is the value smaller than the mapping, we simply go to the next mapping. This is repeated until the number is 0.
