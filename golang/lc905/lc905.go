package lc905

import "sort"

func sortArrayByParity(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		iEven := nums[i]%2 == 0
		jEven := nums[j]%2 == 0
		return iEven && !jEven
	})
	return nums
}
