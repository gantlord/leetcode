package lc15

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	output := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, l, r := -nums[i], i+1, len(nums)-1
		for l < r {
			s := nums[l] + nums[r]
			if s == target {
				triplet := []int{nums[i], nums[l], nums[r]}
				slices.Sort(triplet)
				output = append(output, triplet)
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if s < target {
				l++
			} else {
				r--
			}
		}
	}
	return output
}
