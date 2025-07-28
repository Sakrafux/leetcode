package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}

// Solution
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	init := x
	res := 0
	for x != 0 {
		digit := x % 10
		res = res*10 + digit
		x /= 10
	}

	return res == init
}
