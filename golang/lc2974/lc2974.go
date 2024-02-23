package lc2974

import "sort"

func numberGame(nums []int) []int {
	l := len(nums)
	sort.Ints(nums)
	result := make([]int, l)
	for i := 0; i < l/2; i++ {
		result[i*2] = nums[1]
		result[i*2+1] = nums[0]
		nums = nums[2:]
	}
	return result
}
