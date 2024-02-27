package lc2913

func sumCounts(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			subarray := nums[i : j+1]
			distinct := map[int]bool{}
			for _, e := range subarray {
				distinct[e] = true
			}
			result += len(distinct) * len(distinct)
		}
	}
	return result
}
