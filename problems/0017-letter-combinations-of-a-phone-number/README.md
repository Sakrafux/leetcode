# [17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/)

Given a string containing digits from <code>2-9</code> inclusive, return all possible letter combinations that the number could represent. Return the answer in **any order** .

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
<img alt="" src="https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png" style="width: 300px; height: 243px;">

**Example 1:**

```
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Example 2:**

```
Input: digits = ""
Output: []
```

**Example 3:**

```
Input: digits = "2"
Output: ["a","b","c"]
```

**Constraints:**

- <code>0 <= digits.length <= 4</code>
- <code>digits[i]</code> is a digit in the range <code>['2', '9']</code>.

## Solution

For the problem we first need a hash-map containing all the mappings of digits to letters. Then we iterate over each digit and then further iterate over each possible letter of that digit. The resulting combinations of previous combinations with the new letters overwrite the previous combinations. This continues until we have iterated over all digits.

The time complexity is `O(4^n)`, where `n` is the length of `digits`, as each character has, at maximum, 4 interpretations that are applied to each previous combination.
