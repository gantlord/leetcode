package lc2824

func countPairs(nums []int, target int) int {
	count := 0
	for i, l := range nums {
		for _, r := range nums[i+1:] {
			if l+r < target {
				count++
			}
		}
	}
	return count
}
