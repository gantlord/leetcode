package lc2932

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maximumStrongPairXor(nums []int) int {
	maxXor := math.MinInt
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			ni, nj := nums[i], nums[j]
			if abs(ni-nj) <= min(ni, nj) {
				maxXor = max(maxXor, ni^nj)
			}
		}
	}
	return maxXor
}
