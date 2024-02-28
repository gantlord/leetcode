package lc1827

func minOperations(nums []int) int {
	count := 0
	prev := nums[0]
	for _, n := range nums[1:] {
		diff := n - prev
		offset := 0
		if diff <= 0 {
			offset = 1 - diff
			count += offset
		}
		prev = n + offset
	}
	return count
}
