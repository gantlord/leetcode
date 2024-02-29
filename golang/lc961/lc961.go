package lc961

func repeatedNTimes(nums []int) int {
	seen := map[int]bool{}
	for i := 0; ; i++ {
		n := nums[i]
		if seen[n] {
			return n
		}
		seen[n] = true
	}
}
