# [3541. Find Most Frequent Vowel and Consonant](https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/description/?envType=daily-question&envId=2025-09-13)

You are given a string <code>s</code> consisting of lowercase English letters (<code>'a'</code> to <code>'z'</code>).

Your task is to:

- Find the vowel (one of <code>'a'</code>, <code>'e'</code>, <code>'i'</code>, <code>'o'</code>, or <code>'u'</code>) with the **maximum**  frequency.
- Find the consonant (all other letters excluding vowels) with the **maximum**  frequency.

Return the sum of the two frequencies.

**Note** : If multiple vowels or consonants have the same maximum frequency, you may choose any one of them. If there are no vowels or no consonants in the string, consider their frequency as 0.
The **frequency**  of a letter <code>x</code> is the number of times it occurs in the string.

**Example 1:**

<div class="example-block">
Input: s = "successes"

Output: 6

Explanation:

- The vowels are: <code>'u'</code> (frequency 1), <code>'e'</code> (frequency 2). The maximum frequency is 2.
- The consonants are: <code>'s'</code> (frequency 4), <code>'c'</code> (frequency 2). The maximum frequency is 4.
- The output is <code>2 + 4 = 6</code>.

**Example 2:**

<div class="example-block">
Input: s = "aeiaeia"

Output: 3

Explanation:

- The vowels are: <code>'a'</code> (frequency 3), <code>'e'</code> ( frequency 2), <code>'i'</code> (frequency 2). The maximum frequency is 3.
- There are no consonants in <code>s</code>. Hence, maximum consonant frequency = 0.
- The output is <code>3 + 0 = 3</code>.

**Constraints:**

- <code>1 <= s.length <= 100</code>
- <code>s</code> consists of lowercase English letters only.

## Solution

We can simply solve this problem by iterating the string and keeping track of the number of occurrences of each character
in a map. We need 2 maps to separate vowels and consonants (or treat them differently when using one map). As the range
of characters is limited to lower case letters, we can also use an array effectively as a map. At the end we simply need
to sum up the maximum values from both maps.

The time complexity is `O(n)`.