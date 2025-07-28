package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("42"))            // 42
	fmt.Println(myAtoi(" -042"))         // -42
	fmt.Println(myAtoi("1337c0d3"))      // 1337
	fmt.Println(myAtoi("0-1"))           // 0
	fmt.Println(myAtoi("words and 987")) // 0
	fmt.Println(myAtoi("2147483648"))    // 2147483647
}

// Solution
func myAtoi(s string) int {
	// Handle empty string
	if len(s) == 0 {
		return 0
	}

	i := 0

	// Skip all whitespaces
	for j, char := range s {
		i = j
		if char != ' ' {
			break
		}
	}

	// Handle whitespace-only string
	if s[i] == ' ' {
		return 0
	}

	// Read a sign, if it exists
	sign := 1
	switch s[i] {
	case '-':
		sign = -1
		i++
	case '+':
		i++
	}

	res := 0

	// Iterate over the remaining string
	for ; i < len(s); i++ {
		char := s[i]
		// End if non-numeric character
		if char < '0' || char > '9' {
			break
		}
		// Skip leading zeroes
		if res == 0 && char == '0' {
			continue
		}

		// Add digit to the result
		res = res*10 + int(char-'0')

		// Round the result if necessary
		if signedRes := res * sign; signedRes > math.MaxInt32 {
			return math.MaxInt32
		} else if signedRes < math.MinInt32 {
			return math.MinInt32
		}
	}

	return res * sign
}
