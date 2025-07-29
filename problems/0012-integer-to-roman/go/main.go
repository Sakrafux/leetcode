package main

import (
	"fmt"
)

func main() {
	fmt.Println(intToRoman(3749)) // MMMDCCXLIX
	fmt.Println(intToRoman(58))   // LVIII
	fmt.Println(intToRoman(1994)) // MCMXCIV
}

// Solution
type Mapping struct {
	value int
	roman []byte
}

func intToRoman(num int) string {
	// Create all the possible unique mappings in descending direction
	valueMappings := []Mapping{
		{1000, []byte{'M'}},
		{900, []byte{'C', 'M'}},
		{500, []byte{'D'}},
		{400, []byte{'C', 'D'}},
		{100, []byte{'C'}},
		{90, []byte{'X', 'C'}},
		{50, []byte{'L'}},
		{40, []byte{'X', 'L'}},
		{10, []byte{'X'}},
		{9, []byte{'I', 'X'}},
		{5, []byte{'V'}},
		{4, []byte{'I', 'V'}},
		{1, []byte{'I'}},
	}

	res := make([]byte, 0)

	for i := 0; num > 0; {
		mapping := valueMappings[i]
		// If the number is bigger than the current mapping, then we can
		// substract it and the partial string
		if num >= mapping.value {
			num -= mapping.value
			res = append(res, mapping.roman...)
			// Otherwise, go to the next binding
		} else {
			i++
		}
	}

	return string(res)
}
