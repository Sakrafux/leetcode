# [2785. Sort Vowels in a String](https://leetcode.com/problems/sort-vowels-in-a-string/description/?envType=daily-question&envId=2025-09-11)

Given a **0-indexed**  string <code>s</code>, **permute**  <code>s</code> to get a new string <code>t</code> such that:

- All consonants remain in their original places. More formally, if there is an index <code>i</code> with <code>0 <= i < s.length</code> such that <code>s[i]</code> is a consonant, then <code>t[i] = s[i]</code>.
- The vowels must be sorted in the **nondecreasing**  order of their **ASCII**  values. More formally, for pairs of indices <code>i</code>, <code>j</code> with <code>0 <= i < j < s.length</code> such that <code>s[i]</code> and <code>s[j]</code> are vowels, then <code>t[i]</code> must not have a higher ASCII value than <code>t[j]</code>.

Return the resulting string.

The vowels are <code>'a'</code>, <code>'e'</code>, <code>'i'</code>, <code>'o'</code>, and <code>'u'</code>, and they can appear in lowercase or uppercase. Consonants comprise all letters that are not vowels.

**Example 1:**

```
Input: s = "lEetcOde"
Output: "lEOtcede"
Explanation: 'E', 'O', and 'e' are the vowels in s; 'l', 't', 'c', and 'd' are all consonants. The vowels are sorted according to their ASCII values, and the consonants remain in the same places.
```

**Example 2:**

```
Input: s = "lYmpH"
Output: "lYmpH"
Explanation: There are no vowels in s (all characters in s are consonants), so we return "lYmpH".
```

**Constraints:**

- <code>1 <= s.length <= 10^5</code>
- <code>s</code> consists only of letters of theEnglish alphabetin **uppercase and lowercase** .

## Solution

The solution is rather straightforward. First we copy the original string `s` to `t`, then iterate over the string and 
keep track of each vowel and its location in the string. Now sort only the vowels (not their locations) and use the
original locations to re-insert them, creating the sought-for result.

The time complexity of this is `O(n*log(n))`, where `n = s.length` due to the sorting of the vowels.