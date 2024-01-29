package lc27

func removeElement(nums []int, val int) int {
	outputIndex := 0
	for i := range nums {
		elem := nums[i]
		if elem != val {
			nums[outputIndex] = elem
			outputIndex += 1
		}
	}
	return outputIndex
}
