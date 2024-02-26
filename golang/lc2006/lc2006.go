package lc2006

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countKDifference(nums []int, k int) int {
	count := 0
	for i, numi := range nums {
		for _, numj := range nums[i+1:] {
			if abs(numi-numj) == k {
				count++
			}
		}
	}
	return count
}
