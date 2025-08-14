# [2106. Maximum Fruits Harvested After at Most K Steps](https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/description/?envType=daily-question&envId=2025-08-03)

Fruits are available at some positions on an infinite x-axis. You are given a 2D integer array <code>fruits</code> where <code>fruits[i] = [position<sub>i</sub>, amount<sub>i</sub>]</code> depicts <code>amount<sub>i</sub></code> fruits at the position <code>position<sub>i</sub></code>. <code>fruits</code> is already **sorted** by <code>position<sub>i</sub></code> in **ascending order** , and each <code>position<sub>i</sub></code> is **unique** .

You are also given an integer <code>startPos</code> and an integer <code>k</code>. Initially, you are at the position <code>startPos</code>. From any position, you can either walk to the **left or right** . It takes **one step** to move **one unit** on the x-axis, and you can walk **at most** <code>k</code> steps in total. For every position you reach, you harvest all the fruits at that position, and the fruits will disappear from that position.

Return the **maximum total number** of fruits you can harvest.

**Example 1:**
<img alt="" src="https://assets.leetcode.com/uploads/2021/11/21/1.png" style="width: 472px; height: 115px;">

```
Input: fruits = [[2,8],[6,3],[8,6]], startPos = 5, k = 4
Output: 9
Explanation:
The optimal way is to:
- Move right to position 6 and harvest 3 fruits
- Move right to position 8 and harvest 6 fruits
You moved 3 steps and harvested 3 + 6 = 9 fruits in total.
```

**Example 2:**
<img alt="" src="https://assets.leetcode.com/uploads/2021/11/21/2.png" style="width: 512px; height: 129px;">

```
Input: fruits = [[0,9],[4,1],[5,7],[6,2],[7,4],[10,9]], startPos = 5, k = 4
Output: 14
Explanation:
You can move at most k = 4 steps, so you cannot reach position 0 nor 10.
The optimal way is to:
- Harvest the 7 fruits at the starting position 5
- Move left to position 4 and harvest 1 fruit
- Move right to position 6 and harvest 2 fruits
- Move right to position 7 and harvest 4 fruits
You moved 1 + 3 = 4 steps and harvested 7 + 1 + 2 + 4 = 14 fruits in total.
```

**Example 3:**
<img alt="" src="https://assets.leetcode.com/uploads/2021/11/21/3.png" style="width: 476px; height: 100px;">

```
Input: fruits = [[0,3],[6,4],[8,5]], startPos = 3, k = 2
Output: 0
Explanation:
You can move at most k = 2 steps and cannot reach any position with fruits.
```

**Constraints:**

- <code>1 <= fruits.length <= 10^5</code>
- <code>fruits[i].length == 2</code>
- <code>0 <= startPos, position<sub>i</sub> <= 2 \* 10^5</code>
- <code>position<sub>i-1</sub> < position<sub>i</sub></code> for any <code>i > 0</code>(**0-indexed** )
- <code>1 <= amount<sub>i</sub> <= 10^4</code>
- <code>0 <= k <= 2 \* 10^5</code>

## Solution

The most elegant solution for this problem uses a sliding window. Since the fruits are already sorted ascending by
their position, we can continually move rightwards, while catching up from the left side if we are covering too much
distance. The rightwards expansion adds the number of fruit, while the left-side shrinking removes them. In the worst
case of no close enough fruits, we add and remove all elements resulting in sum 0. For each window we calculate its
distance by the minimum of either going left or right first. Since this uses a sliding window and we only traverse
the array once, this has time complexity `O(n)`.

Another solution would be to use maps and prefix sums. The prefix sum arrays capture the sum of fruits of moving a given 
number of steps in a direction (i.e. 2 prefix sum arrays, left and right). The result is then trying all combinations
of both prefix sums by moving a number of steps in one direction and the remaining in the other. While this technically
also has time complexity `O(n)`, the fixed time operations involved in this are much more numerous thus leading to higher
runtime than the sliding window approach.