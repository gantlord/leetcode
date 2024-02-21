package lc209

import "math"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := nums[0]
	minLen := math.MaxInt

	for l < len(nums) {
		if sum >= target {
			if (r - l + 1) < minLen {
				minLen = r - l + 1
			}
			sum -= nums[l]
			l++
		} else {
			r++
			if r == len(nums) {
				break
			}
			sum += nums[r]
		}
	}
	if minLen == math.MaxInt {
		return 0
	} else {
		return minLen
	}
}
