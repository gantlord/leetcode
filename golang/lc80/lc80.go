package lc80

func removeDuplicates(nums []int) int {
	outputIndex := 0
	var prevElem, prevPrevElem int
	for i, elem := range nums {
		if i == 0 || i == 1 || elem != prevElem || elem != prevPrevElem {
			nums[outputIndex] = elem
			outputIndex += 1
		}
		prevPrevElem = prevElem
		prevElem = elem
	}
	return outputIndex
}
