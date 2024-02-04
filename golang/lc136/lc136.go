package lc136

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	result := nums[len(nums)-1]
	for i := 1; i < len(nums); i += 2 {
		if nums[i-1] != nums[i] {
			result = nums[i-1]
			break
		}
	}
	return result
}
