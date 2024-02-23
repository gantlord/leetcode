package lc2574

import "math"

func leftRightDifference(nums []int) []int {
	result := make([]int, len(nums))
	leftSum := make([]int, len(nums))
	rightSum := make([]int, len(nums))
	lSum, rSum := 0, 0
	for i := 1; i < len(nums); i++ {
		lSum += nums[i-1]
		leftSum[i] = lSum
		rSum += nums[len(nums)-i]
		rightSum[len(nums)-i-1] = rSum
	}
	for i := 0; i < len(nums); i++ {
		result[i] = int(math.Abs(float64(leftSum[i] - rightSum[i])))
	}
	return result
}
