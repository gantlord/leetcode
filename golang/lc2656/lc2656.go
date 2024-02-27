package lc2656

func maximizeSum(nums []int, k int) int {
	maxN := 0
	for _, n := range nums {
		maxN = max(maxN, n)
	}
	return maxN*k + (k*(k-1))/2
}
