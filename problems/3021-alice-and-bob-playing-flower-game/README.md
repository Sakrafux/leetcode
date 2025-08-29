# [3021. Alice and Bob Playing Flower Game](https://leetcode.com/problems/alice-and-bob-playing-flower-game/description/?envType=daily-question&envId=2025-08-29)

Alice and Bob are playing a turn-based game on a field, with two lanes of flowers between them. There are <code>x</code> flowers in the first lane between Alice and Bob, and <code>y</code> flowers in the second lane between them.

<img alt="" src="https://assets.leetcode.com/uploads/2025/08/27/3021.png" style="width: 300px; height: 150px;">

The game proceeds as follows:

- Alice takes the first turn.
- In each turn, a player must choose either one of the laneand pick one flower from that side.
- At the end of the turn, if there are no flowers left at all, the **current**  player captures their opponent and wins the game.

Given two integers, <code>n</code> and <code>m</code>, the task is to compute the number of possible pairs <code>(x, y)</code> that satisfy the conditions:

- Alice must win the game according to the described rules.
- The number of flowers <code>x</code> in the first lane must be in the range <code>[1,n]</code>.
- The number of flowers <code>y</code> in the second lane must be in the range <code>[1,m]</code>.

Return the number of possible pairs <code>(x, y)</code> that satisfy the conditions mentioned in the statement.

**Example 1:**

```
Input: n = 3, m = 2
Output: 3
Explanation: The following pairs satisfy conditions described in the statement: (1,2), (3,2), (2,1).
```

**Example 2:**

```
Input: n = 1, m = 1
Output: 0
Explanation: No pairs satisfy the conditions described in the statement.
```

**Constraints:**

- <code>1 <= n, m <= 10^5</code>

## Solution

We can solve this problem purely using math. It is important to note that a pair can only be valid, if `x` and `y` are
of different parity, i.e. 1 odd and 1 even. As the actual value doesn't matter, but only the parity, we can calculate
the number of odd numbers smaller than `n` times the number of even numbers smaller than `m` and vice versa. We ensure
that, if possible, `n` is odd so we can consider that extra pair. If both `n` and `m` are odd, one needs to ignore that
property anyway, since we need pairs of different parity.

The time complexity is just `O(1)`, since we only do a simple calculation.