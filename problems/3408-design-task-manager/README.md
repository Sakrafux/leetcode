# [3408. Design Task Manager](https://leetcode.com/problems/design-task-manager/description/?envType=daily-question&envId=2025-09-18)

There is a task management system that allows users to manage their tasks, each associated with a priority. The system should efficiently handle adding, modifying, executing, and removing tasks.

Implement the <code>TaskManager</code> class:

-
<code>TaskManager(vector<vector<int>>&amp; tasks)</code> initializes the task manager with a list of user-task-priority triples. Each element in the input list is of the form <code>[userId, taskId, priority]</code>, which adds a task to the specified user with the given priority.

-
<code>void add(int userId, int taskId, int priority)</code> adds a task with the specified <code>taskId</code> and <code>priority</code> to the user with <code>userId</code>. It is **guaranteed**  that <code>taskId</code> does not exist in the system.

-
<code>void edit(int taskId, int newPriority)</code> updates the priority of the existing <code>taskId</code> to <code>newPriority</code>. It is **guaranteed**  that <code>taskId</code> exists in the system.

-
<code>void rmv(int taskId)</code> removes the task identified by <code>taskId</code> from the system. It is **guaranteed**  that <code>taskId</code> exists in the system.

-
<code>int execTop()</code> executes the task with the **highest**  priority across all users. If there are multiple tasks with the same **highest**  priority, execute the one with the highest <code>taskId</code>. After executing, the** ** <code>taskId</code>** ** is **removed**  from the system. Return the <code>userId</code> associated with the executed task. If no tasks are available, return -1.

**Note**  that a user may be assigned multiple tasks.

**Example 1:**

<div class="example-block">
Input:<br>
["TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"]<br>
[[[[1, 101, 10], [2, 102, 20], [3, 103, 15]]], [4, 104, 5], [102, 8], [], [101], [5, 105, 15], []]

Output:<br>
[null, null, null, 3, null, null, 5]

Explanation

TaskManager taskManager = new TaskManager([[1, 101, 10], [2, 102, 20], [3, 103, 15]]); // Initializes with three tasks for Users 1, 2, and 3.<br>
taskManager.add(4, 104, 5); // Adds task 104 with priority 5 for User 4.<br>
taskManager.edit(102, 8); // Updates priority of task 102 to 8.<br>
taskManager.execTop(); // return 3. Executes task 103 for User 3.<br>
taskManager.rmv(101); // Removes task 101 from the system.<br>
taskManager.add(5, 105, 15); // Adds task 105 with priority 15 for User 5.<br>
taskManager.execTop(); // return 5. Executes task 105 for User 5.

**Constraints:**

- <code>1 <= tasks.length <= 10^5</code>
- <code>0 <= userId <= 10^5</code>
- <code>0 <= taskId <= 10^5</code>
- <code>0 <= priority <= 10^9</code>
- <code>0 <= newPriority <= 10^9</code>
- At most <code>2 * 10^5</code> calls will be made in **total**  to <code>add</code>, <code>edit</code>, <code>rmv</code>, and <code>execTop</code> methods.
- The input is generated such that <code>taskId</code> will be valid.

## Solution

The problem itself is not that difficult to understand. We want a priority queue, or a max heap. However, the challenge
lies in the concrete implementation, especially considering we have a bunch of information that belongs together and
regarding more complicated features such as removing an element from the heap. The key lies in a custom heap implementation
using objects that capture all the information including their position in the heap. This allows a map referencing those
objects by task-id to know which element from the heap to remove/edit. Additionally, we need to consider the edit operation
both a removal and addition, instead of a single change operation. 

The time complexity is `O(n*log(n))`, with the construction of the heap dominating. Each subsequent operation is within 
`O(log(n))` due to heap mechanics.