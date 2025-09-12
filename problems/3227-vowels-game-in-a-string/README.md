# [3227. Vowels Game in a String](https://leetcode.com/problems/vowels-game-in-a-string/description/?envType=daily-question&envId=2025-09-12)

Alice and Bob are playing a game on a string.

You are given a string <code>s</code>, Alice and Bob will take turns playing the following game where Alice starts **first** :

- On Alice's turn, she has to remove any **non-empty**  <button type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:r1n:" data-state="closed" class="">substring</button> from <code>s</code> that contains an **odd**  number of vowels.
- On Bob's turn, he has to remove any **non-empty**  <button type="button" aria-haspopup="dialog" aria-expanded="false" aria-controls="radix-:r1o:" data-state="closed" class="">substring</button> from <code>s</code> that contains an **even**  number of vowels.

The first player who cannot make a move on their turn loses the game. We assume that both Alice and Bob play **optimally** .

Return <code>true</code> if Alice wins the game, and <code>false</code> otherwise.

The English vowels are: <code>a</code>, <code>e</code>, <code>i</code>, <code>o</code>, and <code>u</code>.

**Example 1:**

<div class="example-block">
Input: s = "leetcoder"

Output: true

Explanation:<br>
Alice can win the game as follows:

- Alice plays first, she can delete the underlined substring in <code>s = "**leetco** der"</code> which contains 3 vowels. The resulting string is <code>s = "der"</code>.
- Bob plays second, he can delete the underlined substring in <code>s = "**d** er"</code> which contains 0 vowels. The resulting string is <code>s = "er"</code>.
- Alice plays third, she can delete the whole string <code>s = "**er** "</code> which contains 1 vowel.
- Bob plays fourth, since the string is empty, there is no valid play for Bob. So Alice wins the game.

**Example 2:**

<div class="example-block">
Input: s = "bbcd"

Output: false

Explanation:<br>
There is no valid play for Alice in her first turn, so Alice loses the game.

**Constraints:**

- <code>1 <= s.length <= 10^5</code>
- <code>s</code> consists only of lowercase English letters.

## Solution

This problem is more of a brain-teaser than an algorithmic problem. It boils down to the fact that Alice always wins, if
there is at least one vowel in the string. If there is an odd number of vowels, she can take the entire string; if there
is an even number of vowels she can take it in 2 turns. Bob never gets to make a proper move. The only case Alice looses
is if she can't make any valid moves, i.e. there are no vowels at all.

The time complexity is simply `O(n)`, as we iterate over the string to check for vowels.