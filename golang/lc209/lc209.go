package lc209

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := nums[0]
	minLen := 100001

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
	if minLen == 100001 {
		return 0
	} else {
		return minLen
	}
}
