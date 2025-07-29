package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}

// Solution
func romanToInt(s string) int {
	n := len(s)
	charToInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	res := 0

	for i, char := range s {
		// Check the value of the next char
		nextVal := 0
		if i+1 < n {
			nextVal = charToInt[rune(s[i+1])]
		}

		val := charToInt[char]

		// If the next char value is higher, then we have a substraction notation
		if val < nextVal {
			res -= val
		} else {
			res += val
		}
	}

	return res
}
