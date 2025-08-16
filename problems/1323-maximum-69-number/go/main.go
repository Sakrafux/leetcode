package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximum69Number(9669)) // 9969
	fmt.Println(maximum69Number(6669)) // 9669
	fmt.Println(maximum69Number(9996)) // 9999
	fmt.Println(maximum69Number(9999)) // 9999
}

// Solution
func maximum69Number(num int) int {
	digits := make([]int, 0)

	// Extract all digits from the number
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	res := 0
	hasSwapped := false

	// Rebuild the number and change the first occurrence of 6 to 9
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if !hasSwapped && digit == 6 {
			hasSwapped = true
			res = res*10 + 9
		} else {
			res = res*10 + digit
		}
	}

	return res
}
