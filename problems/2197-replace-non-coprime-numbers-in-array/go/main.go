package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(replaceNonCoprimes([]int{6, 4, 3, 2, 7, 6, 2}))                          // [12,7,6]
	fmt.Println(replaceNonCoprimes([]int{2, 2, 1, 1, 3, 3, 3}))                          // [2,1,1,3]
	fmt.Println(replaceNonCoprimes([]int{517, 11, 121, 517, 3, 51, 3, 1887, 5}))         // [5687,1887,5]
	fmt.Println(replaceNonCoprimes([]int{287, 41, 49, 287, 899, 23, 23, 20677, 5, 825})) // [2009,20677,825]
}

// Solution
func replaceNonCoprimes(nums []int) []int {
	result := make([]int, 0)
	// Iterate through all numbers
	for _, num := range nums {
		// Append the last number to the array
		result = append(result, num)
		// We can compare as long as we have more than 1 element
		for len(result) > 1 {
			// Get the last...
			a := result[len(result)-1]
			// ...and the second-to-last number
			b := result[len(result)-2]
			// Calculate their GCD
			g := gcd(a, b)
			if g > 1 {
				// Remove the numbers from the array
				result = result[:len(result)-2]
				// Calculate the LCM using the GCD
				lcm := a / g * b
				// Append the LCM
				result = append(result, lcm)
			} else {
				// Break condition prevents endless looping with no change
				break
			}
		}
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Bad solution - time out
func _replaceNonCoprimes(nums []int) []int {
	factors := make(map[int]map[int]int)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		factors[num] = calculateFactors(num)
	}

	gcds := make(map[[2]int]int)
	lcms := make(map[[2]int]int)

	for i := 1; i < len(nums); {
		aNum, bNum := nums[i-1], nums[i]
		a, b := factors[nums[i-1]], factors[nums[i]]

		gcd := 0
		if gcdValue, ok := gcds[[2]int{aNum, bNum}]; ok {
			gcd = gcdValue
		} else {
			gcd = calculateGcd(a, b)
			gcds[[2]int{aNum, bNum}] = gcd
		}

		if gcd > 1 {
			lcm := 0
			if lcmValue, ok := lcms[[2]int{aNum, bNum}]; ok {
				lcm = lcmValue
			} else {
				lcm = calculateLcm(a, b)
				lcms[[2]int{aNum, bNum}] = lcm
			}

			prevNums := nums
			nums = append(prevNums[:i-1], lcm)
			nums = append(nums, prevNums[i+1:]...)
			if _, ok := factors[lcm]; !ok {
				factors[lcm] = calculateFactors(lcm)
			}
			if i > 1 {
				i--
			}
		} else {
			i++
		}
	}

	return nums
}

func calculateLcm(a, b map[int]int) int {
	keys := make(map[int]bool)
	for k := range a {
		keys[k] = true
	}
	for k := range b {
		keys[k] = true
	}

	lcm := 1
	for k := range keys {
		lcm *= int(math.Pow(float64(k), float64(max(a[k], b[k]))))
	}
	return lcm
}

func calculateGcd(a, b map[int]int) int {
	gcd := 1
	for k, av := range a {
		if bv, ok := b[k]; ok {
			gcd *= min(av, bv) * k
		}
	}
	return gcd
}

func calculateFactors(n int) map[int]int {
	factors := make(map[int]int)

	for i := 2; i <= n; {
		if n%i == 0 {
			factors[i]++
			n /= i
		} else {
			i++
		}
	}

	return factors
}
