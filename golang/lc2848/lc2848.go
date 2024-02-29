package lc2848

func numberOfPoints(nums [][]int) int {
	covered := map[int]bool{}
	for _, c := range nums {
		for i := c[0]; i <= c[1]; i++ {
			covered[i] = true
		}
	}
	return len(covered)
}
