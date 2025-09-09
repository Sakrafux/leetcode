package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(peopleAwareOfSecret(6, 2, 4))      // 5
	fmt.Println(peopleAwareOfSecret(4, 1, 3))      // 6
	fmt.Println(peopleAwareOfSecret(6, 1, 2))      // 2
	fmt.Println(peopleAwareOfSecret(684, 18, 496)) // 653668527
}

// Solution
func peopleAwareOfSecret(n int, delay int, forget int) int {
	// Module constant
	MOD := int(math.Pow10(9) + 7)

	// DP matrix of dimensions n x forget, representing a day i and the number of people knowing the secret for exactly j days
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, forget)
	}

	// Set start
	dp[0][0] = 1
	// Iterate all days
	for i := 1; i < n; i++ {
		// Propagate diagonally, i.e. people who previously knew 2 days now know for 3 days
		for j := 1; j < forget; j++ {
			dp[i][j] = dp[i-1][j-1]
		}
		// Sum all the people who know long enough to share, which generates the new baseline
		for j := delay - 1; j < forget-1; j++ {
			dp[i][0] += dp[i-1][j] % MOD
		}
	}

	// Sum all the aware people on the last day
	aware := 0
	for _, val := range dp[n-1] {
		aware += val
	}

	return aware % MOD
}

func _peopleAwareOfSecret(n int, delay int, forget int) int {
	sharing := make([]int, n+forget)
	forgetting := make([]int, n+forget)
	sharing[0] = 1

	aware := 0

	for i := range n {
		aware = aware + sharing[i] - forgetting[i]
		for j := i + delay; j < i+forget; j++ {
			sharing[j] += sharing[i]
		}
		forgetting[i+forget] += sharing[i]
	}

	return aware % int(math.Pow10(9)+7)
}
