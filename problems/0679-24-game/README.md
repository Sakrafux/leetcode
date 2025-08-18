# [679. 24 Game](https://leetcode.com/problems/24-game/description/?envType=daily-question&envId=2025-08-18)

You are given an integer array <code>cards</code> of length <code>4</code>. You have four cards, each containing a number in the range <code>[1, 9]</code>. You should arrange the numbers on these cards in a mathematical expression using the operators <code>['+', '-', '*', '/']</code> and the parentheses <code>'('</code> and <code>')'</code> to get the value 24.

You are restricted with the following rules:

- The division operator <code>'/'</code> represents real division, not integer division.

- For example, <code>4 / (1 - 2 / 3) = 4 / (1 / 3) = 12</code>.

- Every operation done is between two numbers. In particular, we cannot use <code>'-'</code> as a unary operator.

- For example, if <code>cards = [1, 1, 1, 1]</code>, the expression <code>"-1 - 1 - 1 - 1"</code> is **not allowed** .

- You cannot concatenate numbers together

- For example, if <code>cards = [1, 2, 1, 2]</code>, the expression <code>"12 + 12"</code> is not valid.

Return <code>true</code> if you can get such expression that evaluates to <code>24</code>, and <code>false</code> otherwise.

**Example 1:**

```
Input: cards = [4,1,8,7]
Output: true
Explanation: (8-4) * (7-1) = 24
```

**Example 2:**

```
Input: cards = [1,2,1,2]
Output: false
```

**Constraints:**

- <code>cards.length == 4</code>
- <code>1 <= cards[i] <= 9</code>

## Solution

We can apply recursive depth-first search to brute-force every combination of operations and test them against 24. 
For this we select 2 numbers from a given list, calculate every viable operator (taking care of division-by-0), 
and then recursively try again with a list containing the remaining elements and a single calculation result.
This means that for every step deeper, the list of numbers shortens by 1. If we only have a single number remaining,
then this represents the complete calculation and this we can check against 24.

As the length of our array is fixed, the time complexity collapses into `O(1)`, however, if we were to ignore that fact
it would in fact be `O((n!)³)`. Since each level of the recursion reduces the length of the list by 1, and spawns
an according number of recursion calls, it has complexity `O(n!)`. The nested loops are part of this recursion and
therefore we have `O((n!)³)`.