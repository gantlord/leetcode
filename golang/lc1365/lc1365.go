package lc1365

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		count := 0
		for _, m := range nums {
			if m < n {
				count++
			}
		}
		result[i] = count
	}

	return result
}
