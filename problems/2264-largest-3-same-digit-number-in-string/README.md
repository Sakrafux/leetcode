# [2264. Largest 3-Same-Digit Number in String](https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/?envType=daily-question&envId=2025-08-14)

You are given a string <code>num</code> representing a large integer. An integer is **good**  if it meets the following conditions:

- It is a **substring**  of <code>num</code> with length <code>3</code>.
- It consists of only one unique digit.

Return the **maximum good ** integer as a **string**  or an empty string <code>""</code> if no such integer exists.

Note:

- A **substring**  is a contiguous sequence of characters within a string.
- There may be **leading zeroes**  in <code>num</code> or a good integer.

**Example 1:**

```
Input: num = "6**777** 133339"
Output: "777"
Explanation: There are two distinct good integers: "777" and "333".
"777" is the largest, so we return "777".
```

**Example 2:**

```
Input: num = "23**000** 19"
Output: "000"
Explanation: "000" is the only good integer.
```

**Example 3:**

```
Input: num = "42352338"
Output: ""
Explanation: No substring of length 3 consists of only one unique digit. Therefore, there are no good integers.
```

**Constraints:**

- <code>3 <= num.length <= 1000</code>
- <code>num</code> only consists of digits.

## Solution

We simply iterate over the string and count how often we have seen the current character. At 3 occurrences we check,
whether it is the largest we have seen so far. The result string is then simply generated based on this maximum character.

The time complexity is simply `O(n)`.