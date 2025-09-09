# [2327. Number of People Aware of a Secret](https://leetcode.com/problems/number-of-people-aware-of-a-secret/description/?envType=daily-question&envId=2025-09-09)

On day <code>1</code>, one person discovers a secret.

You are given an integer <code>delay</code>, which means that each person will **share**  the secret with a new person **every day** , starting from <code>delay</code> days after discovering the secret. You are also given an integer <code>forget</code>, which means that each person will **forget**  the secret <code>forget</code> days after discovering it. A person **cannot**  share the secret on the same day they forgot it, or on any day afterwards.

Given an integer <code>n</code>, return the number of people who know the secret at the end of day <code>n</code>. Since the answer may be very large, return it **modulo**  <code>10^9 + 7</code>.

**Example 1:**

```
Input: n = 6, delay = 2, forget = 4
Output: 5
Explanation:
Day 1: Suppose the first person is named A. (1 person)
Day 2: A is the only person who knows the secret. (1 person)
Day 3: A shares the secret with a new person, B. (2 people)
Day 4: A shares the secret with a new person, C. (3 people)
Day 5: A forgets the secret, and B shares the secret with a new person, D. (3 people)
Day 6: B shares the secret with E, and C shares the secret with F. (5 people)
```

**Example 2:**

```
Input: n = 4, delay = 1, forget = 3
Output: 6
Explanation:
Day 1: The first person is named A. (1 person)
Day 2: A shares the secret with B. (2 people)
Day 3: A and B share the secret with 2 new people, C and D. (4 people)
Day 4: A forgets the secret. B, C, and D share the secret with 3 new people. (6 people)
```

**Constraints:**

- <code>2 <= n <= 1000</code>
- <code>1 <= delay < forget <= n</code>

## Solution

The problem itself is not that hard to solve. A simple solution would be to iterate over the days while keeping an array
of the number of people who share or forget the secret on a given day. While iterating, we modify those entries in the
future so they are correct once we reach that day, all while keeping a sum. 

Another more complex approach is to use dynamic programming with a matrix in which each row represents a day and each column
represents a number of days. E.g. `dp[1][1]` would be the number of people who on the 2. day know the secret exactly for
2 days (current day included). Now for each day we need to 1. propagate the people who knew on the day before (i.e. diagonally),
and 2. calculate who is newly informed by summing all the people of the previous day who are now eligible to share the secret.
Then we only need to sum the number of people on the last day.

In both cases the time complexity is `O(n*m)` with `m = forget`.

However, the true challenge of this problem is not the theoretical solution but rather the fact that we reach way too large
numbers, which leads to integer overflow and thus wrong results. This necessitates using the module operation during the 
summation part of a new day for dynamic programming, which doesn't distort the results. The simple solution doesn't allow
for that.