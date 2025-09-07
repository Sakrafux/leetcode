package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumZero(5)) // [-7,-1,1,3,4]
	fmt.Println(sumZero(3)) // [-1,0,1]
	fmt.Println(sumZero(1)) // [0]
}

// Solution
func sumZero(n int) []int {
	res := make([]int, 0, n)

	// Add symmetric values
	for i := 1; i <= n/2; i++ {
		res = append(res, i, -i)
	}

	// Add 0 if odd number of elements are required
	if n%2 == 1 {
		res = append(res, 0)
	}

	return res
}
