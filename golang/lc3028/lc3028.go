package lc3028

func returnToBoundaryCount(nums []int) int {
	pos, count := 0, 0
	for _, n := range nums {
		pos += n
		if pos == 0 {
			count++
		}
	}
	return count
}
