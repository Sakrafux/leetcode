# [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/description/)

Given <code>n</code> pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

**Example 2:**

```
Input: n = 1
Output: ["()"]
```

**Constraints:**

- <code>1 <= n <= 8</code>

## Solution

The method to solve this problem is dynamic programming. We keep a map of all the valid strings by the number of parantheses as key. We now iteratively build the next layer based on the previous until we reach the target number of parantheses.

In each iteration, we already know that each pattern of the previous solution wrapped with a pair of parantheses is part of the current solution. Additionally, every combination of patterns with that sum up to the current number of parantheses is also part of the solution. We just use a hash map to ensure the uniqueness of the combinations and afterwards append them to our current solution.

The result is available in the map with the given target number.

The time complexity is `O(n^4)`, since we are iterating over `n` once nested, i.e. `O(n²)`, but also trying all combinations of 2 pattern lists that are each of `O(n)`, so `O(n²)` inside `O(n²)`.
