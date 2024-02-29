package lc2089

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	result := []int{}
	for i, n := range nums {
		if n == target {
			result = append(result, i)
		}
	}
	return result
}
