package lc1480

func runningSum(nums []int) []int {
	result := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		result[i] = sum
	}
	return result
}
