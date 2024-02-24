package lc1389

func createTargetArray(nums []int, index []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		p := index[i]
		result = append(result[:p], append([]int{nums[i]}, result[p:]...)...)
	}
	return result
}
