package lc2733

import "math"

func findNonMinOrMax(nums []int) int {
	minN, maxN, res := math.MaxInt, math.MinInt, -1
	for _, n := range nums {
		minN = min(minN, n)
		maxN = max(maxN, n)
	}
	for _, n := range nums {
		if n != minN && n != maxN {
			res = n
		}
	}
	return res
}
