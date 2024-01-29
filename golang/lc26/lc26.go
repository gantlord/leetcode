package lc26

func removeDuplicates(nums []int) int {
	outputIndex := 0
	for i, elem := range nums {
		nums[outputIndex] = elem
		if i == 0 || elem != nums[outputIndex-1] {
			nums[outputIndex] = elem
			outputIndex += 1
		}
	}
	return outputIndex
}
