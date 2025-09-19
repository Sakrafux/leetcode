# [3484. Design Spreadsheet](https://leetcode.com/problems/design-spreadsheet/description/?envType=daily-question&envId=2025-09-19)

A spreadsheet is a grid with 26 columns (labeled from <code>'A'</code> to <code>'Z'</code>) and a given number of <code>rows</code>. Each cell in the spreadsheet can hold an integer value between 0 and 10^5.

Implement the <code>Spreadsheet</code> class:

- <code>Spreadsheet(int rows)</code> Initializes a spreadsheet with 26 columns (labeled <code>'A'</code> to <code>'Z'</code>) and the specified number of rows. All cells are initially set to 0.
- <code>void setCell(String cell, int value)</code> Sets the value of the specified <code>cell</code>. The cell reference is provided in the format <code>"AX"</code> (e.g., <code>"A1"</code>, <code>"B10"</code>), where the letter represents the column (from <code>'A'</code> to <code>'Z'</code>) and the number represents a **1-indexed**  row.
- <code>void resetCell(String cell)</code> Resets the specified cell to 0.
- <code>int getValue(String formula)</code> Evaluates a formula of the form <code>"=X+Y"</code>, where <code>X</code> and <code>Y</code> are **either**  cell references or non-negative integers, and returns the computed sum.

**Note:**  If <code>getValue</code> references a cell that has not been explicitly set using <code>setCell</code>, its value is considered 0.

**Example 1:**

<div class="example-block">
Input:<br>
["Spreadsheet", "getValue", "setCell", "getValue", "setCell", "getValue", "resetCell", "getValue"]<br>
[[3], ["=5+7"], ["A1", 10], ["=A1+6"], ["B2", 15], ["=A1+B2"], ["A1"], ["=A1+B2"]]

Output:<br>
[null, 12, null, 16, null, 25, null, 15]

Explanation

Spreadsheet spreadsheet = new Spreadsheet(3); // Initializes a spreadsheet with 3 rows and 26 columns<br data-end="321" data-start="318">
spreadsheet.getValue("=5+7"); // returns 12 (5+7)<br data-end="373" data-start="370">
spreadsheet.setCell("A1", 10); // sets A1 to 10<br data-end="423" data-start="420">
spreadsheet.getValue("=A1+6"); // returns 16 (10+6)<br data-end="477" data-start="474">
spreadsheet.setCell("B2", 15); // sets B2 to 15<br data-end="527" data-start="524">
spreadsheet.getValue("=A1+B2"); // returns 25 (10+15)<br data-end="583" data-start="580">
spreadsheet.resetCell("A1"); // resets A1 to 0<br data-end="634" data-start="631">
spreadsheet.getValue("=A1+B2"); // returns 15 (0+15)

**Constraints:**

- <code>1 <= rows <= 10^3</code>
- <code>0 <= value <= 10^5</code>
- The formula is always in the format <code>"=X+Y"</code>, where <code>X</code> and <code>Y</code> are either valid cell references or **non-negative**  integers with values less than or equal to <code>10^5</code>.
- Each cell reference consists of a capital letter from <code>'A'</code> to <code>'Z'</code> followed by a row number between <code>1</code> and <code>rows</code>.
- At most <code>10^4</code> calls will be made in **total**  to <code>setCell</code>, <code>resetCell</code>, and <code>getValue</code>.

## Solution

The solution to this problem is very straightforward, because the input is so heavily restricted. Thus, we don't need
to particularly care about the spreadsheet/matrix semantics of the data and can simply treat it as a hash map with default
values. For getting the values we simply need to parse the string and decide whether it is a cell reference or integer literal.

The time complexity is `O(1)` for all operations, including instantiation.