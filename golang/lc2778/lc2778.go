package lc2778

func sumOfSquares(nums []int) int {
	result := 0
	for i, n := range nums {
		if len(nums)%(i+1) == 0 {
			result += n * n
		}
	}
	return result
}
