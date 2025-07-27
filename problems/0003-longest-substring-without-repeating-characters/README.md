# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** , and each of their nodes contains a single digit. Add the two numbers and return the sumas a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example 1:**
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg" style="width: 483px; height: 342px;">

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

**Constraints:**

- The number of nodes in each linked list is in the range <code>[1, 100]</code>.
- <code>0 <= Node.val <= 9</code>
- It is guaranteed that the list represents a number that does not have leading zeros.

## Solution

To solve the problem we can use a sliding window that represents a duplicate-free substring at all times. As we iterate over the string, we keep track of the last index we saw a specific character using a map. If we see a character again, it possibly means that we have reached the end of the duplicate-free substring and need to move the start of our sliding window.

We need to take to not move if the last time a character was seen was outside the sliding window, i.e. the suggested new window start is behind the current window start.

Then we just need to update the map with the new index of the seen character and check, whether the new length is longer.

### Other attempts

A brute force attempt just creates all possible substrings and then checks their eligability and length. Though this takes too long for this task.

A slightly more sophisticated attempt makes use of a similar duplicate detection using a map, but it does so for each substring. I.e. we don't create the substrings in advance but build them as we traverse the array and check using duplicate chars, whether they are still eligable. While this approach works, it is rather slow.
