# [36. Valid Sudoku](https://leetcode.com/problems/valid-sudoku/description/?envType=daily-question&envId=2025-08-30)

Determine if a<code>9 x 9</code> Sudoku boardis valid.Only the filled cells need to be validated**according to the following rules** :

- Each rowmust contain thedigits<code>1-9</code> without repetition.
- Each column must contain the digits<code>1-9</code>without repetition.
- Each of the nine<code>3 x 3</code> sub-boxes of the grid must contain the digits<code>1-9</code>without repetition.

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentionedrules.

**Example 1:**
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png" style="height: 250px; width: 250px;">

```
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
```

**Example 2:**

```
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the **5**  in the top left corner being modified to **8** . Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```

**Constraints:**

- <code>board.length == 9</code>
- <code>board[i].length == 9</code>
- <code>board[i][j]</code> is a digit <code>1-9</code> or <code>'.'</code>.

## Solution

The solve this problem, we simply need to iterate the board while keeping the previously encountered symbols in memory,
on a row, column, and 3x3 block basis. Hash Maps are fit for this purpose. If we find no duplicates with the maps, then
we can conclude that a given board is valid.

The time complexity is actually `O(1)`, since the dimensions of the board are fixed and hash map access is static.