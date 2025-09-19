package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	spreadsheet := Constructor(3)               // Initializes a spreadsheet with 3 rows and 26 columns
	fmt.Println(spreadsheet.GetValue("=5+7"))   // returns 12 (5+7)
	spreadsheet.SetCell("A1", 10)               // sets A1 to 10
	fmt.Println(spreadsheet.GetValue("=A1+6"))  // returns 16 (10+6)
	spreadsheet.SetCell("B2", 15)               // sets B2 to 15
	fmt.Println(spreadsheet.GetValue("=A1+B2")) // returns 25 (10+15)
	spreadsheet.ResetCell("A1")                 // resets A1 to 0
	fmt.Println(spreadsheet.GetValue("=A1+B2")) // returns 15 (0+15)
}

// Solution
type Spreadsheet struct {
	cells map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{make(map[string]int)}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cells[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.cells[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	parts := strings.Split(formula[1:], "+")
	v1Str := parts[0]
	v2Str := parts[1]

	v1, err := strconv.Atoi(v1Str)
	if err != nil {
		v1 = this.cells[v1Str]
	}

	v2, err := strconv.Atoi(v2Str)
	if err != nil {
		v2 = this.cells[v2Str]
	}

	return v1 + v2
}
