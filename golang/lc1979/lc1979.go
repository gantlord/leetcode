package lc1979

import "math"

func findGCD(nums []int) int {
	small, large := math.MaxInt32, -math.MaxInt32
	for _, n := range nums {
		small = min(small, n)
		large = max(large, n)
	}
	gcd := 0
	for i := 1; i <= small; i++ {
		if small%i == 0 && large%i == 0 {
			gcd = i
		}
	}
	return gcd
}
