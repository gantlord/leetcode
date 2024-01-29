package lc169

import "sort"

func majorityElement(nums []int) int {
	limit := len(nums) / 2
	sort.Ints(nums)
	recordHolder := nums[0]
	recordCount := 1
	prospectiveRecordHolder := nums[0]
	prospectiveRecordCount := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == prospectiveRecordHolder {
			prospectiveRecordCount++
		} else {
			prospectiveRecordHolder = nums[i]
			prospectiveRecordCount = 1
		}
		if prospectiveRecordCount > recordCount {
			recordHolder = prospectiveRecordHolder
			recordCount = prospectiveRecordCount
		}
		if recordCount > limit {
			return recordHolder
		}
	}
	return recordHolder
}
