# [1792. Maximum Average Pass Ratio](https://leetcode.com/problems/maximum-average-pass-ratio/description/?envType=daily-question&envId=2025-09-01)

There is a school that has classes of students and each class will be having a final exam. You are given a 2D integer array <code>classes</code>, where <code>classes[i] = [pass<sub>i</sub>, total<sub>i</sub>]</code>. You know beforehand that in the <code>i^th</code> class, there are <code>total<sub>i</sub></code> total students, but only <code>pass<sub>i</sub></code> number of students will pass the exam.

You are also given an integer <code>extraStudents</code>. There are another <code>extraStudents</code> brilliant students that are **guaranteed**  to pass the exam of any class they are assigned to. You want to assign each of the <code>extraStudents</code> students to a class in a way that **maximizes**  the **average**  pass ratio across **all**  the classes.

The **pass ratio**  of a class is equal to the number of students of the class that will pass the exam divided by the total number of students of the class. The **average pass ratio**  is the sum of pass ratios of all the classes divided by the number of the classes.

Return the **maximum**  possible average pass ratio after assigning the <code>extraStudents</code> students. Answers within <code>10^-5</code> of the actual answer will be accepted.

**Example 1:**

```
Input: classes = [[1,2],[3,5],[2,2]], <code>extraStudents</code> = 2
Output: 0.78333
Explanation: You can assign the two extra students to the first class. The average pass ratio will be equal to (3/4 + 3/5 + 2/2) / 3 = 0.78333.
```

**Example 2:**

```
Input: classes = [[2,4],[3,9],[4,5],[2,10]], <code>extraStudents</code> = 4
Output: 0.53485
```

**Constraints:**

- <code>1 <= classes.length <= 10^5</code>
- <code>classes[i].length == 2</code>
- <code>1 <= pass<sub>i</sub> <= total<sub>i</sub> <= 10^5</code>
- <code>1 <= extraStudents <= 10^5</code>

## Solution

To solve this problem we need to make use of a max-heap, which stores the classes, but orders them according to their
ratio increase, i.e. how much their ratio would increase if one excellent student was added. This ensures that the head
element always provides the greatest increase to the average ratio. Since the amount of elements don't change dynamically,
we don't need to bother with removing/adding to the heap, but only heapifying up during the creation of the heap, and 
heapifying down the head element after adding a student.

The time complexity for this problem is `O(max(n*log(n), m*log(n)))`, since both heapifying up or down are `O(log(n))`.
They are executed `n`-times (number of classes) for the creation, and `m`-times (number of extra students) during modification.