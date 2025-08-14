# [2561. Rearranging Fruits](https://leetcode.com/problems/rearranging-fruits/description/?envType=daily-question&envId=2025-08-02)

You have two fruit baskets containing <code>n</code> fruits each. You are given two **0-indexed** integer arrays <code>basket1</code> and <code>basket2</code> representing the cost of fruit in each basket. You want to make both baskets **equal** . To do so, you can use the following operation as many times as you want:

- Chose two indices <code>i</code> and <code>j</code>, and swap the <code>i<font size="1">th</code>fruit of <code>basket1</code> with the <code>j<font size="1">th</code>fruit of <code>basket2</code>.
- The cost of the swap is <code>min(basket1[i],basket2[j])</code>.

Two baskets are considered equal if sorting them according to the fruit cost makes them exactly the same baskets.

Return the minimum cost to make both the baskets equal or <code>-1</code> if impossible.

**Example 1:**

```
Input: basket1 = [4,2,2,2], basket2 = [1,4,1,2]
Output: 1
Explanation: Swap index 1 of basket1 with index 0 of basket2, which has cost 1. Now basket1 = [4,1,2,2] and basket2 = [2,4,1,2]. Rearranging both the arrays makes them equal.
```

**Example 2:**

```
Input: basket1 = [2,3,4,1], basket2 = [3,2,5,1]
Output: -1
Explanation: It can be shown that it is impossible to make both the baskets equal.
```

**Constraints:**

- <code>basket1.length == basket2.length</code>
- <code>1 <= basket1.length <= 10^5</code>
- <code>1 <= basket1[i],basket2[i]<= 10^9</code>

## Solution

We need to make use of 2 facts: 1. the total number of fruits to split per type and the number of swaps must be even,
because otherwise there is no solution anyway, and 2. instead of swapping 2 elements directly we can instead do 2
swaps using the smallest element as a helper.

Now the first part consists of counting the occurrences in each basket and, based on that, establish the elements that
need to be swapped. The second part concerns itself with optimizing the necessary swaps.

The most complex part of this is the sorting and thus we have time complexity `O(n*log(n))`.