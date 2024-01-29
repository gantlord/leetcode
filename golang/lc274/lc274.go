package lc274

func hIndex(citations []int) int {
	maxIndex := 0
	for i := 1; i <= len(citations); i++ {
		count := 0
		for _, elem := range citations {
			if elem >= i {
				count++
			}
		}
		if count < i {
			return max(0, i-1)
		}

		maxIndex++

	}
	return maxIndex
}
